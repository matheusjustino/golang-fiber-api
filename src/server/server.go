package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/matheusjustino/golang-fiber-api/src/routes"
)

type Server struct {
	port   string
	server *fiber.App
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: fiber.New(),
	}
}

func (s *Server) Run() {
	s.server.Use(cors.New())
	s.server.Use(logger.New())

	routes.SetupRoutes(s.server)

	err := s.server.Listen(":" + s.port)
	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
	log.Println("Server is running at port: " + s.port)
}
