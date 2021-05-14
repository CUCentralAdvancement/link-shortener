package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var getenv = os.Getenv

const (
	xForwardedProtoHeader = "x-forwarded-proto"
	goEnviron             = "GO_ENV"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := fiber.New()

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./pages/index.html")
	})

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	err := app.Listen(":" + port)
	if err != nil {
		return
	}
}
