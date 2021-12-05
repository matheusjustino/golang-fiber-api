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
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (*AuthService) Login(data *models.LoginModel) (*string, *fiber.Error) {
	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user_collection := database.UsersCollection()

	filter := bson.M{
		"email": data.Email,
	}
	findOptions := options.FindOne()

	user := new(schemas.User)

	err := user_collection.FindOne(context, filter, findOptions).Decode(&user)
	if err != nil {
		log.Println(err.Error())
		if err == mongo.ErrNoDocuments {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid credentials")
		} else {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
		}
	}

	validPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if validPassword != nil {
		log.Println(validPassword.Error())
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid credentials")
	}

	jwtService := NewJwtService()
	token, generateTokenErr := jwtService.GenerateToken(user.ID.Hex())
	if generateTokenErr != nil {
		return nil, fiber.NewError(generateTokenErr.Code, generateTokenErr.Message)
	}

	return token, nil
}

func (*AuthService) Register(data *schemas.User) (*string, *fiber.Error) {
	if err := data.Prepare(); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusBadRequest, "Failed while preparing to insert user")
	}

	context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user_collection := database.UsersCollection()

	result, err := user_collection.InsertOne(context, data)
	if err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusBadRequest, "Failed while insert user")
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()

	return &insertedId, nil
}
