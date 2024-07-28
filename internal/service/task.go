package service

import (
	"context"
	"project-management/internal/domain/task"
	interfaces "project-management/internal/repository/interface"
	services "project-management/internal/service/interface"
)

type TaskService struct {
	taskRepository interfaces.TaskRepository
}

func NewTaskService(repository interfaces.TaskRepository) services.TaskService {
	return &TaskService{
		taskRepository: repository,
	}
}

func (ts *TaskService) CreateTask(ctx context.Context, req task.Request) (id string, err error) {
	data := task.Entity{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      req.Status,
		UserID:      task.ParseID(req.UserID),
		ProjectID:   task.ParseID(req.ProjectID),
		StartDate:   task.ParseDate(req.StartDate),
		EndDate:     task.ParseDate(req.EndDate),
	}
	id, err = ts.taskRepository.Create(ctx, data)
	return
}

func (ts *TaskService) ListTasks(ctx context.Context) (res []task.Response, err error) {
	data, err := ts.taskRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	res = task.ParseFromEntities(data)
	return
}

func (ts *TaskService) GetTask(ctx context.Context, id string) (res task.Response, err error) {
	data, err := ts.taskRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = task.ParseFromEntity(data)
	return
}

func (ts *TaskService) DeleteTask(ctx context.Context, id string) (err error) {
	err = ts.taskRepository.Delete(ctx, id)
	return
}

func (ts *TaskService) UpdateTask(ctx context.Context, id string, req task.Request) (err error) {
	data := task.Entity{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		UserID:      task.ParseID(req.UserID),
		ProjectID:   task.ParseID(req.ProjectID),
		StartDate:   task.ParseDate(req.StartDate),
		EndDate:     task.ParseDate(req.EndDate),
	}
	err = ts.taskRepository.Update(ctx, id, data)
	return
}

func (ts *TaskService) SearchTask(ctx context.Context, filter, value string) (res []task.Response, err error) {
	if !task.IsValidFilter(filter) || value == "" {
		err = task.ErrorInvalidSearch
		return
	}

	data, err := ts.taskRepository.Search(ctx, filter, value)
	if err != nil {
		return nil, err
	}
	res = task.ParseFromEntities(data)
	return
}
