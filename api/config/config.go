package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Database struct {
	Url      string `env:"APP_DB_URL, required"`
	Port     int    `env:"APP_DB_PORT, required"`
	Name     string `env:"APP_DB_NAME, required"`
	Username string `env:"APP_DB_USERNAME, required"`
	Password string `env:"APP_DB_PASSWORD, required"`
	SslMode  string `env:"APP_DB_SSL_MODE, required" split_words:"true"`
}

type AppConfig struct {
	Host          string `env:"APP_HOST, required"`
	Port          int    `env:"APP_PORT, required"`
	DB            Database
	ProjectSecret string `env:"APP_PROJECT_SECRET, required" split_words:"true"`
}

func GetConfig() AppConfig {
	_ = godotenv.Overload()

	var appConf AppConfig

	err := envconfig.Process("app", &appConf)
	if err != nil {
		fmt.Print("Cannot load env")
	}

	return appConf
}
