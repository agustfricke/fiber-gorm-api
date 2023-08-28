package main

import (
	"log"

	"github.com/agustfricke/fiber-crud-api/database"
	"github.com/agustfricke/fiber-crud-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    database.ConnectDB()

    app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

    router.SetupRoutes(app)

    log.Fatal(app.Listen(":8000"))
}
