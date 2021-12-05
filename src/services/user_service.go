package services

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusjustino/golang-fiber-api/src/database"
	"github.com/matheusjustino/golang-fiber-api/src/database/schemas"
	"github.com/matheusjustino/golang-fiber-api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct{}

func (*UserService) FindAll() (*[]schemas.User, *fiber.Error) {
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user_collection := database.UsersCollection()
	filter := bson.M{}            // Opcionais
	findOptions := options.Find() // Opcionais

	var users []schemas.User

	cursor, err := user_collection.Find(context, filter, findOptions)
	if err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusNotFound, "Error whiling get users")
	}

	for cursor.Next(context) {
		var user schemas.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	err = cursor.Close(context)
	if err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error whiling decode users")
	}

	return &users, nil
}

func (*UserService) FindOne(id primitive.ObjectID) (*schemas.User, *fiber.Error) {
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user_collection := database.UsersCollection()

	var user schemas.User

	findOptions := options.FindOne()
	filter := bson.M{"_id": id}

	err := user_collection.FindOne(context, filter, findOptions).Decode(&user)
	if err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return &user, nil
}

func (*UserService) UpdateUser(id primitive.ObjectID, data *models.UpdateUserModel) (*schemas.User, *fiber.Error) {
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user_collection := database.UsersCollection()

	var user schemas.User

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":      data.Name,
			"email":     data.Email,
			"password":  data.Password,
			"update_at": time.Now(),
		},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	err := user_collection.FindOneAndUpdate(context, filter, update, &opt).Decode(&user)
	if err != nil {
		log.Println("error:", err.Error())
		if err == mongo.ErrNoDocuments {
			return nil, fiber.NewError(fiber.StatusNotFound, "Failed while update user")
		} else {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
	}

	return &user, nil
}

func (*UserService) DeleteUser(id primitive.ObjectID) *fiber.Error {
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user_collection := database.UsersCollection()

	filter := bson.M{"_id": id}
	deleteOptions := options.FindOneAndDelete()

	err := user_collection.FindOneAndDelete(context, filter, deleteOptions).Err()
	if err != nil {
		log.Println(err.Error())
		if err == mongo.ErrNoDocuments {
			return fiber.NewError(fiber.StatusNotFound, "Failed while delete user")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
	}

	return nil
}
