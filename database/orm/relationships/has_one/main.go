package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	category := Category{
		Name: "Game",
	}
	db.Create(&category)

	productFifa := Product{
		Name:       "Fifa 24",
		CategoryID: category.ID,
		Price:      500,
	}
	db.Create(&productFifa)

	var products []Product
	db.Preload("Category").Find(&products)

	for _, p := range products {
		fmt.Printf(
			"ID: %d - Name: %s - CategoryID: %d\n",
			p.ID,
			p.Name,
			p.CategoryID,
		)
	}
}
