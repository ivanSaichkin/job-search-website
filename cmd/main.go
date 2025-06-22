package main

import (
	"job-search/go-fiber/config"
	"job-search/go-fiber/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	dbConfig := config.NewDatabaseConfig()
	log.Println(dbConfig)
	app := fiber.New()
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
