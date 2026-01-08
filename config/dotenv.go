package config

import "github.com/joho/godotenv"

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}
}
