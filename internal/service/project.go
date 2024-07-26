package service

import (
	"context"
	"project-management/internal/domain/project"
	"project-management/internal/repository"
)

type ProjectService struct {
	projectRepository *repository.ProjectRepository
}

func NewProjectService(repository *repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepository: repository,
	}
}

func (ps *ProjectService) CreateProject(ctx context.Context, req project.Request) (id string, err error) {
	data := project.Entity{
		Title:       req.Title,
		Description: req.Description,
		ManagerID:   req.ManagerID,
	}

	id, err = ps.projectRepository.Create(ctx, data)
	if err != nil {
		return
	}

	return
}
func (ps *ProjectService) ListProjects(ctx context.Context) (res []project.Response, err error) {
	data, err := ps.projectRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	res = project.ParseFromEntities(data)
	return
}
func (ps *ProjectService) GetProject()    {}
func (ps *ProjectService) DeleteProject() {}
func (ps *ProjectService) UpdateProject() {}
