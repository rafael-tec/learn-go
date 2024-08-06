package usecase_test

import (
	"context"
	"database/sql"
	"testing"
	"uow/internal/repository"
	"uow/internal/usecase"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func Test_CreateCourse_Success(t *testing.T) {
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbConn.Exec("DROP TABLE IF EXISTS courses;")
	dbConn.Exec("DROP TABLE IF EXISTS categories;")

	dbConn.Exec("CREATE TABLE IF NOT EXISTS categories (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbConn.Exec("CREATE TABLE IF NOT EXISTS courses (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id int NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	input := usecase.InputUseCase{
		CategoryName:     "Music",
		CourseName:       "Beat maker",
		CourseCategoryID: 2,
	}

	useCase := usecase.NewCreateCourseUseCase(
		repository.NewCourseRepository(dbConn),
		repository.NewCategoryRepository(dbConn),
	)

	ctx := context.Background()
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
