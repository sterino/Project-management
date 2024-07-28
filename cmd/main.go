package main

import (
	"log"
	"os"
	"project-management/internal/config"
	"project-management/internal/di"

	_ "project-management/docs"
)

// @title Project Management
// @version 1.0
// @description API Server for Project Management App
// @BasePath /api
func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal(diErr)
	} else {
		server.Run(infoLog, errorLog)
	}
}
