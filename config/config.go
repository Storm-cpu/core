package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DbPsn string

	Port         int
	ReadTimeout  int
	WriteTimeout int
}

func LoadConfig() Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %+v", err)
	}

	dbPsn := os.Getenv("DB_PSN")

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	readTimeout, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	writeTimeout, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))

	config := Configuration{
		DbPsn: dbPsn,

		Port:         port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return config
}
