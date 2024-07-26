package service

import (
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

//
//func (ts *TaskService) CreateTask(ctx context.Context, req task.Request) (err error) {
//	return nil
//}
//
//func (ts *TaskService) ListTasks(ctx context.Context) (err error) {
//	return nil
//}
//
//func (ts *TaskService) GetTask(ctx context.Context) (err error) {
//	return nil
//}
//
//func (ts *TaskService) DeleteTask(ctx context.Context) (err error) {
//	return nil
//}
//
//func (ts *TaskService) UpdateTask(ctx context.Context) (err error) {
//	return nil
//}
