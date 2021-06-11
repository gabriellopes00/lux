package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DB_URI = ""
	PORT   = 0
)

func LoadVars() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	DB_URI = fmt.Sprintf("%s?sslmode=disable", os.Getenv("DB_URI"))
}
