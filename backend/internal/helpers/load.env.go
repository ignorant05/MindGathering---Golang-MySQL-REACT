package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while loading environment variables, \n%v", err)
	}
}
