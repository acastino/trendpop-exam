package main

import (
	"os"

	"github.com/joho/godotenv"
)

func loadVars() {
	godotenv.Load()
}

func portNumByEnv() string {
	sysPort := os.Getenv("PORT")
	if sysPort != "" {
		return ":" + sysPort
	}
	return ":8001"
}
