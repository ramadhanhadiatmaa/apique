package routes

import (
	"apique/controllers"
	"apique/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	queue := r.Group("/api")

	queue.Get("/", middlewares.AuthMiddleware, controllers.Index)
	queue.Get("/:id", middlewares.AuthMiddleware, controllers.Show)
	queue.Post("/", middlewares.AuthMiddleware, controllers.Create)
	queue.Put("/:id", middlewares.AuthMiddleware, controllers.Update)
	queue.Delete("/:id", middlewares.AuthMiddleware, controllers.Delete)
}