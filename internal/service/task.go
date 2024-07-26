package service

import (
	"context"
	"project-management/internal/repository"
)

type TaskService struct {
	taskRepository *repository.TaskRepository
}

func NewTaskService(repository *repository.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: repository,
	}
}

func (ts *TaskService) CreateTask(ctx context.Context) {

}

func (ts *TaskService) ListTasks(ctx context.Context) {

}

func (ts *TaskService) GetTask(ctx context.Context) {

}

func (ts *TaskService) DeleteTask(ctx context.Context) {

}

func (ts *TaskService) UpdateTask(ctx context.Context) {

}
