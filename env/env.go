package env

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env", ".env.local")

	if err != nil {
		log.Println(err, "- ignore this when using Docker")
	}
}
