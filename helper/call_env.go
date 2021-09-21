package helper

import (
	"github.com/joho/godotenv"
	"log"
)

func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return
}
