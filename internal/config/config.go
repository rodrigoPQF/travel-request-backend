package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	POSTGRES_USER    string 
	POSTGRES_PASSWORD    string 
	POSTGRES_DB    string 
	POSTGRES_HOST    string 
	PORT     string 
}

func InitAppConfig() error {
	_ = godotenv.Load()

	return validateEmptyEnv()
}

func GetEnvs() EnvVars{
	env := EnvVars{}
	env.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	env.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	env.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	env.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	env.PORT = os.Getenv("PORT")
	return env
}

func validateEmptyEnv() error {
	required := []string{"PORT", "POSTGRES_PASSWORD", "POSTGRES_DB", "POSTGRES_HOST", "POSTGRES_USER"}

	for _, v := range required {
		if value := os.Getenv(v); value == "" {
			 return fmt.Errorf("error empty environment %s", v)
		}
	}
	return nil
}
