package interfaces

import (
	"context"
	"project-management/internal/domain/task"
)

type TaskRepository interface {
	Create(ctx context.Context, entity task.Entity) (id string, err error)
	List(ctx context.Context) (res []task.Entity, err error)
	Get(ctx context.Context, id string) (res task.Entity, err error)
	Delete(ctx context.Context, id string) (err error)
	Update(ctx context.Context, id string, entity task.Entity) (err error)
	Search(ctx context.Context, filter, value string) (res []task.Entity, err error)
}
