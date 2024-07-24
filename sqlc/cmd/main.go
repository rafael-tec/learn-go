package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/rafael-tec/learn-go/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	ctx := context.Background()
	_, err = queries.CreateCategory(
		ctx,
		db.CreateCategoryParams{
			ID:          uuid.New().String(),
			Name:        "Backend",
			Description: sql.NullString{String: "Backend", Valid: true},
		},
	)

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category)
	}
}
