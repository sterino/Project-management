package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (cfg Config, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		cfg = Config{
			DBHost:     os.Getenv("DBHost"),
			DBPort:     os.Getenv("DBPort"),
			DBUser:     os.Getenv("DBUser"),
			DBPassword: os.Getenv("DBPassword"),
			DBName:     os.Getenv("DBName"),
		}
		return
	}

	if err = envconfig.Process("DB", &cfg); err != nil {
		return
	}

	return
}
