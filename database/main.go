package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	product := NewProduct("Playstation 5", 5.000)
	err = saveProduct(db, *product)
	if err != nil {
		panic(err)
	}

	newPrice := 4.500
	product.Price = newPrice
	err = updateProduct(db, *product)
	if err != nil {
		panic(err)
	}

	resultByID, err := fetchProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Select by id result: %v\n\n", resultByID)

	resultAll, err := fetchProducts(db)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Select all result: %v\n\n", resultAll)
}

func saveProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("UPDATE products SET name=?, price=? WHERE id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func fetchProduct(db *sql.DB, productId string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id=?")
	if err != nil {
		return &Product{}, err
	}

	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(productId).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return &Product{}, err
	}

	return &p, nil
}

func fetchProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return []Product{}, err
	}

	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return []Product{}, err
		}

		products = append(products, p)
	}

	return products, nil
}
