package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusjustino/golang-fiber-api/src/models"
	"github.com/matheusjustino/golang-fiber-api/src/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	userService *services.UserService
}

func (c *UserController) FindAll(ctx *fiber.Ctx) error {
	users, findAllErr := c.userService.FindAll()
	if findAllErr != nil {
		return ctx.Status(findAllErr.Code).JSON(fiber.Map{
			"success": false,
			"message": findAllErr.Message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}

func (c *UserController) FindOne(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id := idParam[:len(idParam)-3]

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	user, findOneErr := c.userService.FindOne(objectId)
	if findOneErr != nil {
		return ctx.Status(findOneErr.Code).JSON(fiber.Map{
			"success": false,
			"message": findOneErr.Message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id := idParam[:len(idParam)-3]

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create objectId from param id",
			"error":   err.Error(),
		})
	}

	user := new(models.UpdateUserModel)

	if err := ctx.BodyParser(user); err != nil {
		log.Println(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err.Error(),
		})
	}

	_, updateUserErr := c.userService.UpdateUser(objectId, user)
	if updateUserErr != nil {
		return ctx.Status(updateUserErr.Code).JSON(fiber.Map{
			"success": false,
			"message": updateUserErr.Message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"user":    id,
	})
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id := idParam[:len(idParam)-3]

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create objectId from param id",
			"error":   err.Error(),
		})
	}

	deleteUserErr := c.userService.DeleteUser(objectId)
	if deleteUserErr != nil {
		return ctx.Status(deleteUserErr.Code).JSON(fiber.Map{
			"success": false,
			"message": deleteUserErr.Message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"user":    id,
	})
}
