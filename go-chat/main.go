package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file")
	}

	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   os.Getenv("APP_ID"),
		Key:     os.Getenv("TOKEN_KEY"),
		Secret:  os.Getenv("TOKEN_SECRET"),
		Cluster: "eu",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)

		return c.JSON([]string{})
	})

	log.Fatal(app.Listen(":8000"))
}
