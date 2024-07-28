package interfaces

import (
	"context"
	"project-management/internal/domain/task"
)

type TaskService interface {
	CreateTask(ctx context.Context, req task.Request) (id string, err error)
	ListTasks(ctx context.Context) (res []task.Response, err error)
	GetTask(ctx context.Context, id string) (res task.Response, err error)
	DeleteTask(ctx context.Context, id string) (err error)
	UpdateTask(ctx context.Context, id string, req task.Request) (err error)
	SearchTask(ctx context.Context, filter, value string) (res []task.Response, err error)
}
