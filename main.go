package main

import (
	"github.com/JimmyTanapon/Hexagonal/adepters"

	"github.com/JimmyTanapon/Hexagonal/core"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("fail to connect database!")
	}
	orderRepo := adepters.NewGormOrderRepository(db)
	orderService := core.NewOrdeService(orderRepo)
	orderHandler := adepters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)
	db.AutoMigrate(&core.Order{})
	app.Listen(":8000")
}
