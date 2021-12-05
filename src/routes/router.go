package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/matheusjustino/golang-fiber-api/src/middleware"
)

var auth_middleware middleware.JwtMiddleware

func SetupRoutes(app *fiber.App) {
	// Monitorar o consumo de recursos
	app.Get("/dashboard", auth_middleware.JWTGuard(), monitor.New())

	api := app.Group("/api/v1")

	AuthRoutes(api.Group("/auth"))
	UserRoutes(api.Group("/users", auth_middleware.JWTGuard()))
}
