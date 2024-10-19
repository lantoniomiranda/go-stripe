package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	StripeSecretKey string
}

var AppConfig Config

func LoadConfig() {
	loadEnv()
	AppConfig = Config{
		StripeSecretKey: getEnv("STRIPE_SECRET_KEY", ""),
	}

	if AppConfig.StripeSecretKey == "" {
		log.Fatal("Stripe secret key is missing")
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
