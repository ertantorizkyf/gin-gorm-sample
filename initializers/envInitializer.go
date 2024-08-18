package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	// LOAD .ENV FILE
	err := godotenv.Load()

	if err != nil {
		log.Printf("[ERR] Error loading .env file")
	} else {
		log.Printf("[INFO] .env file loaded")
	}
}
