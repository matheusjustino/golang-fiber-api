package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusjustino/golang-fiber-api/src/database/schemas"
	"github.com/matheusjustino/golang-fiber-api/src/models"
	"github.com/matheusjustino/golang-fiber-api/src/services"
)

type AuthController struct {
	authService *services.AuthService
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	data := new(models.LoginModel)

	err := ctx.BodyParser(&data)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	token, loginErr := c.authService.Login(data)
	if loginErr != nil {
		return ctx.Status(loginErr.Code).JSON(fiber.Map{
			"success": false,
			"message": loginErr.Message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	user := schemas.NewUser()

	if err := ctx.BodyParser(&user); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"data":    err.Error(),
		})
	}

	result, err := c.authService.Register(user)
	if err != nil {
		log.Println(err.Error())
		return ctx.Status(err.Code).JSON(fiber.Map{
			"success": false,
			"message": err.Message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User inserted successfully",
		"data":    result,
	})
}
