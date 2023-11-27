package main

import (
	"os"

	"github.com/joho/godotenv"
)

func loadVars() {
	godotenv.Load()
}

func portNumByEnv() string {
	if os.Getenv("ENV") == "production" {
		return ":80"
	} else {
		return ":8001"
	}
}
