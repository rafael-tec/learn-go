package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{
		Name: "Game",
	}
	db.Create(&category)

	product := Product{
		Name:       "Fifa 24",
		CategoryID: category.ID,
		Price:      500,
	}
	db.Create(&product)

	var saved Product
	db.Last(&saved)

	db.Create(&SerialNumber{
		Number:    uuid.NewString(),
		ProductID: saved.ID,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Printf(
			"ID: %d - Name: %s - CategoryID: %d - SerialNumber: %s\n",
			p.ID,
			p.Name,
			p.CategoryID,
			p.SerialNumber.Number,
		)
	}
}
