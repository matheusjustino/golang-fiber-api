package main

import (
	"github.com/matheusjustino/golang-fiber-api/src/config"
	"github.com/matheusjustino/golang-fiber-api/src/database"
	"github.com/matheusjustino/golang-fiber-api/src/server"
)

func main() {
	config.LoadEnv()
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
