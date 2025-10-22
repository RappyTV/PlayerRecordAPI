package env

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Failed to load .env file: ", err)
	} else {
		log.Println("Environment variables loaded from .env file")
	}
}
