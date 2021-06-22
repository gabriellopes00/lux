package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DB_USER      = ""
	DB_PASS      = ""
	DB_PORT      = 0
	DB_HOST      = ""
	DB_NAME      = ""
	DB_SSL_MODE  = ""
	DB_TIME_ZONE = ""

	PORT = 0
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}

	DB_PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_SSL_MODE = os.Getenv("DB_SSL_MODE")
	DB_TIME_ZONE = os.Getenv("DB_TIME_ZONE")

	return nil
}
