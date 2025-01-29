package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load environment variables here
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}