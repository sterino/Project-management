package interfaces

import (
	"context"
	"project-management/internal/domain/task"
	"project-management/internal/domain/user"
)

type UserService interface {
	CreateUser(ctx context.Context, req user.Request) (id string, err error)
	ListUsers(ctx context.Context) (res []user.Response, err error)
	GetUser(ctx context.Context, id string) (res user.Response, err error)
	DeleteUser(ctx context.Context, id string) (err error)
	UpdateUser(ctx context.Context, id string, req user.Request) (err error)
	SearchUser(ctx context.Context, filter, value string) (res []user.Response, err error)
	GetUserTasks(ctx context.Context, UserID string) (res []task.Response, err error)
}
