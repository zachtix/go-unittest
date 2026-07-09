package main

import (
	// "zachtix/fiber"
	"fmt"
	"zachtix/gorm"
)

func main() {
	// app := fiber.Setup()
	// app.Listen(":8000")

	db := gorm.InitializePostgres()
	// db := gorm.InitializeSQLite()
	err := gorm.AddUser(db, "John Doe", "john.doe@mail.com", 30)
	if err != nil {
		fmt.Println("error", err)
	}
}
