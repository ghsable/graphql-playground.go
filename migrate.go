package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string
	Name     string
	IsActive bool
}

func main() {

	db := dbInit()

	db.AutoMigrate(&User{})

	insert(db)

	fmt.Println("DONE.")
}

func dbInit() *gorm.DB {
	dsn := "dev.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func insert(db *gorm.DB) {
	users := []User{
		{
			Email:    "a@xxx.com",
			Name:     "name-a",
			IsActive: true,
		},
		{
			Email:    "b@xxx.com",
			Name:     "name-b",
			IsActive: true,
		},
		{
			Email:    "c@xxx.com",
			Name:     "name-c",
			IsActive: true,
		},
		{
			Email:    "d@xxx.com",
			Name:     "name-d",
			IsActive: false,
		},
		{
			Email:    "e@xxx.com",
			Name:     "name-e",
			IsActive: false,
		},
	}
	result := db.Create(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("inserted:", result.RowsAffected)
}
