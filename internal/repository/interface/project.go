package interfaces

import (
	"context"
	"project-management/internal/domain/project"
)

type ProjectRepository interface {
	Create(ctx context.Context, data project.Entity) (id string, err error)
	List(ctx context.Context) (projects []project.Entity, err error)
	Get(ctx context.Context, id string) (dest project.Entity, err error)
	Delete(ctx context.Context, id string) (err error)
	Update(ctx context.Context, id string, data project.Entity) (err error)
	Search(ctx context.Context, filter, value string) (dest []project.Entity, err error)
}
