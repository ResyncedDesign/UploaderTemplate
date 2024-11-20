package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetServerPort() string {
	return os.Getenv("PORT") // Port to run the server on (pretty self-explanatory)
}

func GetR2URL() string {
	return os.Getenv("R2_URL") // URL to your R2 storage service (e.g. https://r2.interrupted.me)
}
