package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fn RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *UOW) error) error
	CommitOrRollback() error
	UnRegister(name string)
}

type UOW struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUOW(ctx context.Context, db *sql.DB) *UOW {
	return &UOW{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *UOW) Register(name string, fn RepositoryFactory) {
	u.Repositories[name] = fn
}

func (u *UOW) UnRegister(name string) {
	delete(u.Repositories, name)
}

func (u *UOW) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}

		u.Tx = tx
	}

	rep := u.Repositories[name](u.Tx)
	return rep, nil
}

func (u *UOW) Do(ctx context.Context, fn func(UOW *UOW) error) error {
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	u.Tx = tx
	err = fn(u)
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}

		return err
	}

	return u.CommitOrRollback()
}

func (u *UOW) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}
	u.Tx = nil
	return nil
}

func (u *UOW) Rollback() error {
	if u.Tx == nil {
		return errors.New("no transaction to rollback")
	}

	err := u.Tx.Rollback()
	if err != nil {
		return err
	}

	u.Tx = nil
	return nil
}
