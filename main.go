package main

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// PostgreSQL connection string
	dsn := "user=postgres password=Welcome@1234 dbname=postgres host=localhost port=5433 sslmode=disable"

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		log.Println("Connected to the database successfully.")
	}

	
	db.AutoMigrate(&Product{})

	
	db.Create(&Product{Code: "D42", Price: 100})

	
	var product Product
	result := db.First(&product, 1)

	
	if result.Error != nil {
		log.Println("Error fetching product:", result.Error)
	} else {
		log.Printf("Product found: %v\n", product)
	}
}
