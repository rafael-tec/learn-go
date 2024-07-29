package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/rafael-tec/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTX(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %w, origin error: %w", errRb, err)
		}

		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(
	ctx context.Context,
	courseParams CourseParams,
	categoryParams CategoryParams,
) error {
	err := c.callTX(ctx, func(q *db.Queries) error {
		var err error
		_, err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          categoryParams.ID,
			Name:        categoryParams.Name,
			Description: categoryParams.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          courseParams.ID,
			Name:        courseParams.Name,
			Description: courseParams.Description,
			Price:       courseParams.Price,
			CategoryID:  categoryParams.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func main() {
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	courseDB := NewCourseDB(dbConn)

	ctx := context.Background()
	courses, err := courseDB.ListCoursesJoinCategory(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("List courses")
	for _, course := range courses {
		fmt.Printf("%+v \n", course)
	}

	courseParams := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: sql.NullString{String: "Learn Golang for backend", Valid: true},
		Price:       99.65,
	}

	categoryParams := CategoryParams{
		ID:   uuid.New().String(),
		Name: "Technology",
		Description: sql.NullString{
			String: "Digital Design, Machine Learning, Programing",
			Valid:  true,
		},
	}

	err = courseDB.CreateCourseAndCategory(ctx, courseParams, categoryParams)
	if err != nil {
		panic(err)
	}
}
