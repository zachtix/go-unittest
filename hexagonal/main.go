package main

import (
	"zachtix/hexagonal/adapters"
	"zachtix/hexagonal/core"

	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&core.Order{})

	// Set up the core service and adapters
	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	// Define routes
	app.Post("/order", orderHandler.CreateOrder)

	// Start the server
	app.Listen(":8000")
}
