package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	rootDir, _ := os.Getwd()
	_, err := os.Stat(filepath.Join(rootDir, ".env"))

	// Carrega o arquivo .env se ele existir
	if err == nil {
		err = godotenv.Load()

		if err != nil {
			log.Fatalf("Error loading .env file")
		} else {
			log.Println("Env file loaded")
		}
	}
}
