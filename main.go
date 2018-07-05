package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product Grom entity
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatalf("Error creating DB connection: %v", err)
	}

	defer db.Close()

	// Manage Schema
	db.AutoMigrate(&Product{})

	// Insert
	db.Create(&Product{
		Code:  "L123",
		Price: 100,
	})

	// Select
	var myProduct Product
	db.First(&myProduct, 1)
	fmt.Printf("Find by ID: %v", myProduct)

	db.First(&myProduct, "code = ?", "L123")
	fmt.Printf("Find by code: %v", myProduct)

	// Update
	db.Model(&myProduct).Update("Price", 200)
	db.First(&myProduct, 1)
	fmt.Printf("After update: %v", myProduct)

	// Delte
	db.Delete(&myProduct)
}
