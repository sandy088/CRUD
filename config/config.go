package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI      string
	MongoDatabase string
}

func LoadConfig() *Config {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Println("Error loading .env file", err)
	}

	return &Config{
		MongoURI:      os.Getenv("MONGO_URI"),
		MongoDatabase: os.Getenv("MONGO_DATABASE"),
	}
}
