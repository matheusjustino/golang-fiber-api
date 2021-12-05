package services

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type jwtService struct {
	secretKey string
	issure    string
}

type Claim struct {
	Sum string `json:"sum"`
	jwt.StandardClaims
}

func NewJwtService() *jwtService {
	return &jwtService{
		secretKey: os.Getenv("SECRET"),
		issure:    "golang-fiber-api",
	}
}

func (s *jwtService) GenerateToken(id string) (*string, *fiber.Error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed while generate token")
	}

	return &signedToken, nil
}
