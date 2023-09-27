package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Some errorcode occured. Err: %s", err)
	}
}
