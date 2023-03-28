package main

import (
  "fmt"

  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {

  db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&Product{})

  // Create
  //db.Create(&Product{Code: "D42", Price: 100})

  fmt.Println("Finish")
}
