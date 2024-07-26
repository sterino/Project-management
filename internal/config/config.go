package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type Config struct {
	DBHost     string `envconfig:"DBHost"`
	DBPort     string `envconfig:"DBPort"`
	DBUser     string `envconfig:"DBUser"`
	DBPassword string `envconfig:"DBPassword"`
	DBName     string `envconfig:"DBName"`
}

func LoadConfig() (cfg Config, err error) {

	if err = envconfig.Process("DB", &cfg); err != nil {
		return
	}
	root, err := os.Getwd()
	if err != nil {
		return
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		cfg.DBHost = os.Getenv("DBHost")
		cfg.DBPort = os.Getenv("DBPort")
		cfg.DBUser = os.Getenv("DBUser")
		cfg.DBPassword = os.Getenv("DBPassword")
		cfg.DBName = os.Getenv("DBName")

		return cfg, nil
	}

	return
}
