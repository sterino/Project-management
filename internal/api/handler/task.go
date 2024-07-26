package handler

import (
	"project-management/internal/service"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: service,
	}
}
