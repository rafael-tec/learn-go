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
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
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

	productFifa := Product{
		Name:       "Fifa 24",
		CategoryID: category.ID,
		Price:      500,
	}
	db.Create(&productFifa)

	db.Create(&SerialNumber{
		Number:    uuid.NewString(),
		ProductID: productFifa.ID,
	})

	categoryNew := Category{
		Name: "Vehicle",
	}
	db.Create(&categoryNew)

	productGTA := Product{
		Name:       "GTA 6",
		CategoryID: categoryNew.ID,
		Price:      500,
	}
	db.Create(&productGTA)

	db.Create(&SerialNumber{
		Number:    uuid.NewString(),
		ProductID: productGTA.ID,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name + ":")
		for _, p := range c.Products {
			fmt.Printf("   ProductName: %s - SerialNumber: %s\n", p.Name, p.SerialNumber.Number)
		}
	}
}
