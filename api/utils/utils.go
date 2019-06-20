package utils

import (
	env "github.com/joho/godotenv"
) //env is like an alias

// LoadEnv should load the env file
func LoadEnv() {
	env.Load()
}
