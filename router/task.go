package router

import (
	"github.com/agustfricke/fiber-crud-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/api/get/tasks/", handlers.GetTasks)
    app.Post("/api/post/tasks/", handlers.CreateTask)
    app.Get("/api/get/task/:id", handlers.GetSoloTask)
    app.Put("/api/put/task/:id", handlers.UpdateTask)
    app.Delete("/api/delete/task/:id", handlers.DeleteTask)
}
