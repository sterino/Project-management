//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	_ "github.com/lib/pq"
	http "project-management/internal/api"
	"project-management/internal/api/handler"
	"project-management/internal/config"
	"project-management/internal/db"
	"project-management/internal/repository"
	"project-management/internal/service"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewProjectRepository,
		service.NewProjectService,
		handler.NewProjectHandler,
		handler.NewTaskHandler,
		handler.NewUserHandler,
		repository.NewUserRepository,
		repository.NewTaskRepository,
		service.NewUserService,
		service.NewTaskService,
		http.NewServer,
	)
	return &http.Server{}, nil
}
