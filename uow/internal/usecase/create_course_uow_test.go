package usecase_test

import (
	"context"
	"database/sql"
	"testing"
	"uow/internal/db"
	"uow/internal/repository"
	"uow/internal/usecase"
	"uow/pkg/uow"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func Test_CreateCourseUOW_Success(t *testing.T) {
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbConn.Exec("DROP TABLE IF EXISTS courses;")
	dbConn.Exec("DROP TABLE IF EXISTS categories;")

	dbConn.Exec("CREATE TABLE IF NOT EXISTS categories (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbConn.Exec("CREATE TABLE IF NOT EXISTS courses (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id int NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	ctx := context.Background()
	uow := uow.NewUOW(ctx, dbConn)
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbConn)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("CourseCategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbConn)
		repo.Queries = db.New(tx)
		return repo
	})

	useCase := usecase.NewCreateCourseUseCaseUow(uow)

	input := usecase.InputUseCase{
		CategoryName:     "Music",
		CourseName:       "Beat maker",
		CourseCategoryID: 2,
	}
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
