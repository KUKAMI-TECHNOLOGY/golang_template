package routers

import (
	"github.com/gofiber/fiber/v2"
	"main.go/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	// Correctly reference the handler function without calling it
	v1.Post("/", handler.CreateUser)
	v1.Get("/", handler.GetAllUser)
	v1.Get("/:id" , handler.GetSingleUser)
	v1.Delete("/:id" , handler.DeleteUserByID)
	v1.Put("/:id", handler.UpdateUser)

}
