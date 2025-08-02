package main

import (
	"job-search/go-fiber/config"
	"job-search/go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()

	log.SetLevel(log.Level(logConfig.Level))
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
