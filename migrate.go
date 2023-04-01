package main

import (
  "fmt"
  "log"

  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type User struct {
  gorm.Model

  Email string
  Name string
  IsActive bool
}

func main() {

  db := dbInit()

  // Migrate the schema
  db.AutoMigrate(&User{})

  // Create
  //db.Create(&Product{Code: "D42", Price: 100})
  inserts(db)

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

func inserts(db *gorm.DB) {
	users := []User{
		{
			Email:       "a@xxx.com",
			Name:        "name-a",
			IsActive:    true,
		},
		{
			Email:       "b@xxx.com",
			Name:        "name-b",
			IsActive:    true,
		},
		{
			Email:       "c@xxx.com",
			Name:        "name-c",
			IsActive:    true,
		},
	}
	result := db.Create(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}
