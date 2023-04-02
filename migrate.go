package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	Email    string
	IsActive bool
}

func main() {
	{
		db := ConnectDB()

		db.AutoMigrate(&User{})

		InsertData(db)
	}

	fmt.Println("DONE.")
}

func ConnectDB() *gorm.DB {
	dsn := "dev.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func InsertData(db *gorm.DB) {
	users := []User{
		{
			Name:     "name-a",
			Email:    "a@xxx.com",
			IsActive: true,
		},
		{
			Name:     "name-b",
			Email:    "b@xxx.com",
			IsActive: true,
		},
		{
			Name:     "name-c",
			Email:    "c@xxx.com",
			IsActive: true,
		},
		{
			Name:     "name-d",
			Email:    "d@xxx.com",
			IsActive: false,
		},
		{
			Name:     "name-e",
			Email:    "e@xxx.com",
			IsActive: false,
		},
	}
	data := &users
	if result := db.Create(data); result.Error != nil {
		log.Fatal(result.Error)
	} else {
		fmt.Println("inserted:", result.RowsAffected)
	}
}
