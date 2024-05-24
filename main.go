package main

import (
	"fmt"

	"github.com/JimmyTanapon/Hexagonal/adepters"

	"github.com/JimmyTanapon/Hexagonal/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5433         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	app := fiber.New()
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database!")
	}
	db.AutoMigrate(&core.Order{})
	orderRepo := adepters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adepters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)
	db.AutoMigrate(&core.Order{})
	app.Listen(":8000")
}
