package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusjustino/golang-fiber-api/src/controllers"
)

var auth_controller controllers.AuthController

func AuthRoutes(route fiber.Router) {
	route.Post("/login", auth_controller.Login)
	route.Post("/register", auth_controller.Register)
}
