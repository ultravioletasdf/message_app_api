package main

import (
	"message_app_api/routes"
	"message_app_api/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := utils.GetConfig()
	app := fiber.New()

	app.Use(compress.New(), helmet.New(), logger.New())

	routes.Add(app)

	app.Listen(config.ListenAddress, fiber.ListenConfig{EnablePrefork: true})
}
