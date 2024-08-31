package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return &Config{}, err
	}

	config := &Config{
		DBHost:     getEnv("DATABASE_HOST", "localhost"),
		DBPort:     getEnv("DATABASE_PORT", "5432"),
		DBUser:     getEnv("DATABASE_USER", "user"),
		DBPassword: getEnv("DATABASE_PASSWORD", "password"),
		DBName:     getEnv("DATABASE_NAME", "dbname"),
		ServerPort: getEnv("DATABASE_PORT", "8080"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}
