package env

import (
	"os"
	"strconv"

	"log"

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

	TOKEN_PUBLIC_KEY      = ""
	TOKEN_PRIVATE_KEY     = ""
	TOKEN_EXPIRATION_TIME = 0

	PORT = 0
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Panicln(err)
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_SSL_MODE = os.Getenv("DB_SSL_MODE")
	DB_TIME_ZONE = os.Getenv("DB_TIME_ZONE")
	DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Panicln(err)
	}

	TOKEN_PRIVATE_KEY = os.Getenv("TOKEN_PRIVATE_KEY")
	TOKEN_PUBLIC_KEY = os.Getenv("TOKEN_PUBLIC_KEY")
	TOKEN_EXPIRATION_TIME, err = strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_TIME"))
	if err != nil {
		log.Panicln(err)
	}

}
