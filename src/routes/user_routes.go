package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusjustino/golang-fiber-api/src/controllers"
)

var user_controller controllers.UserController

func UserRoutes(route fiber.Router) {
	route.Get("", user_controller.FindAll)
	route.Get("/:id", user_controller.FindOne)
	route.Put("/:id", user_controller.UpdateUser)
	route.Delete("/:id", user_controller.DeleteUser)
}
