package routes

import (
	"github.com/gofiber/fiber/v2"
	"user-api-go/internal/handler"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	app.Post("/users", h.Create)
	app.Get("/users", h.List)
	app.Get("/users/:id", h.Get)
	app.Put("/users/:id", h.Update)
	app.Delete("/users/:id", h.Delete)
}
