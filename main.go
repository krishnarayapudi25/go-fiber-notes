package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/krishnarayapudi25/go-fiber-notes/database"
    "github.com/krishnarayapudi25/go-fiber-notes/router"
     
)

func main() {
    // Start a new fiber app
    app := fiber.New()

    database.ConnectDB()

    router.SetupRoutes(app)

    app.Get("/", func(c *fiber.Ctx) error {
        err := c.SendString("And the API is UP!")
        return err
    })

    // Listen on PORT 300
    app.Listen(":3000")
}