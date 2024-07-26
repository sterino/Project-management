package handler

import (
	"github.com/gin-gonic/gin"
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

func (th *TaskHandler) CreateTask(c *gin.Context) {

}

func (th *TaskHandler) ListTasks(c *gin.Context) {

}
func (th *TaskHandler) GetTask(c *gin.Context) {

}
func (th *TaskHandler) UpdateTask(c *gin.Context) {

}
func (th *TaskHandler) DeleteTask(c *gin.Context) {

}
