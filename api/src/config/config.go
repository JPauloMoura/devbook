package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DbConnectionString = ""
	ApiPort            = 0
)

func Load() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load envs", err)
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println("used API_PORT default 3001")
		ApiPort = 3001
	}

	DbConnectionString = os.Getenv("DB_CONNECTION_STRING")
	if DbConnectionString == "" {
		log.Fatal("DB_CONNECTION_STRING is empty")
	}
}
