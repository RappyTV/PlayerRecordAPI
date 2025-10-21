package env

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Failed to load .env file", err)
	}

	log.Println("Environment variables loaded from .env file")
}
