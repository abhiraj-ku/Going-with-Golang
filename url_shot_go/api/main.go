package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoute(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	//server instance of the fiber app
	app := fiber.New()

	// logger of the server
	app.Use(logger.New())

	// call the setup route

	setupRoute(app)

	// starting the server

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
