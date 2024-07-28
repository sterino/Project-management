// project-management/internal/service/interface/project.go
package interfaces

import (
	"context"
	"project-management/internal/domain/project"
	"project-management/internal/domain/task"
)

type ProjectService interface {
	CreateProject(ctx context.Context, req project.Request) (id string, err error)
	ListProjects(ctx context.Context) (res []project.Response, err error)
	GetProject(ctx context.Context, id string) (res project.Response, err error)
	DeleteProject(ctx context.Context, id string) (err error)
	UpdateProject(ctx context.Context, id string, req project.Request) (err error)
	SearchProject(ctx context.Context, filter, value string) (res []project.Response, err error)
	GetProjectTasks(ctx context.Context, projectID string) (res []task.Response, err error)
}
