package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        int
	DBUser        string
	DBPassword    string
	DBName        string
	SecretKey     string
	TokenIssuer   string
	TokenAudience string
}

func LoadConfig() (*Config, error) {
	envPaths := []string{".env", "../.env", "../../.env"}
	var loadErr error
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			// Successfully loaded
			loadErr = nil
			break
		} else {
			loadErr = err
		}
	}

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        dbPort,
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		SecretKey:     os.Getenv("SECRET_KEY"),
		TokenIssuer:   os.Getenv("TOKEN_ISSUER"),
		TokenAudience: os.Getenv("TOKEN_AUDIENCE"),
	}, loadErr
}
