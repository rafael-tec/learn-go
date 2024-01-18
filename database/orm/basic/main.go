package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Catalog struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	CategoryID int
	Category
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Catalog{}, &Category{})

	db.Create(&Catalog{
		Name: "Default",
	})

	electronicsCategory := &Category{
		Name: "Category Electronics",
	}

	db.Create(electronicsCategory)

	electronicsCatalog := &Catalog{
		Name:       "Electronics on sale",
		CategoryID: electronicsCategory.ID,
	}

	insert(db, electronicsCatalog)

	vehicleCategory := &Category{
		Name: "Category Vehicle",
	}

	db.Create(vehicleCategory)

	carCatalog := Catalog{
		Name:       "Car on sale",
		CategoryID: vehicleCategory.ID,
	}

	motorcycleCatalog := Catalog{
		Name:       "Motorcycle on sale",
		CategoryID: vehicleCategory.ID,
	}

	insertBatch(db, &[]Catalog{carCatalog, motorcycleCatalog})

	catalogByID := selectByID(db, 1)
	fmt.Printf("SelectByID: \n%v\n", catalogByID)

	catalogByName := selectByName(db, "Car")
	fmt.Printf("\nSelectByName: \n%v\n", catalogByName)

	catalogs := selectAll(db)
	fmt.Printf("\nSelectAll: \n%v\n", catalogs)

	catalogsByLimit := selectAllWithLimit(db, 5)
	fmt.Printf("\nSelectAllByLimit: \n%v\n", catalogsByLimit)

	catalogsByPagination := selectAllWithLimitAndOffset(db, 3, 4)
	fmt.Printf("\nSelectAllByPagination: \n%v\n", catalogsByPagination)

	catalogsByLike := selectByLike(db, "%a%")
	fmt.Printf("\nSelectByLike: \n%v\n", catalogsByLike)

	catalogByID.Name = "New " + catalogByID.Name
	update(db, catalogByID)
	newCatalog := selectByID(db, catalogByID.ID)
	fmt.Printf("\nUpdate: \n%v\n", newCatalog)

	delete(db, newCatalog.ID)

	catalogsWithJoin := selectAllWithJoin(db)
	fmt.Print("\nSelectWithJoin: \n")

	for _, c := range catalogsWithJoin {
		fmt.Printf(
			"ID: %d\n - Name: %s - CategoryID: %d - CreatedAt: %s",
			c.ID,
			c.Name,
			c.CategoryID,
			c.CreatedAt,
		)
	}
}

func insert(db *gorm.DB, record *Catalog) {
	db.Create(record)
}

func insertBatch(db *gorm.DB, records *[]Catalog) {
	db.Create(records)
}

func selectByID(db *gorm.DB, catalogID int) Catalog {
	var catalog Catalog
	db.First(&catalog, catalogID)
	return catalog
}

func selectByName(db *gorm.DB, catalogName string) Catalog {
	var catalog Catalog
	db.First(&catalog, "name = ?", catalogName)
	return catalog
}

func selectAll(db *gorm.DB) []Catalog {
	var catalogs []Catalog
	db.Find(&catalogs)
	return catalogs
}

func selectAllWithLimit(db *gorm.DB, limit int) []Catalog {
	var catalogs []Catalog
	db.Limit(limit).Find(&catalogs)
	return catalogs
}

func selectAllWithLimitAndOffset(db *gorm.DB, limit int, offset int) []Catalog {
	var catalogs []Catalog
	db.Limit(limit).Offset(offset).Find(&catalogs)
	return catalogs
}

func selectByLike(db *gorm.DB, expression string) []Catalog {
	var catalogs []Catalog
	db.Where("name LIKE ?", expression).Find(&catalogs)
	return catalogs
}

func update(db *gorm.DB, newCatalog Catalog) {
	db.Save(&newCatalog)
}

func delete(db *gorm.DB, catalogID int) {
	db.Where("id = ?", catalogID).Delete(&Catalog{ID: catalogID})
}

func selectAllWithJoin(db *gorm.DB) []Catalog {
	var catalogs []Catalog
	db.Preload("Category").Find(&catalogs)
	return catalogs
}
