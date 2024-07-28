package service

import (
	"context"
	"errors"
	"project-management/internal/domain/project"
	"project-management/internal/domain/task"
	"project-management/internal/domain/user"
	interfaces "project-management/internal/repository/interface"
	services "project-management/internal/service/interface"
)

type ProjectService struct {
	projectRepository interfaces.ProjectRepository
	taskRepository    interfaces.TaskRepository
}

func NewProjectService(projectRepository interfaces.ProjectRepository, taskRepostory interfaces.TaskRepository) services.ProjectService {
	return &ProjectService{
		projectRepository: projectRepository,
		taskRepository:    taskRepostory,
	}
}

func (ps *ProjectService) CreateProject(ctx context.Context, req project.Request) (id string, err error) {
	data := project.Entity{
		Title:       req.Title,
		Description: req.Description,
		ManagerID:   project.ParseID(req.ManagerID),
		StartDate:   project.ParseDate(req.StartDate),
		EndDate:     project.ParseDate(req.EndDate),
	}
	id, err = ps.projectRepository.Create(ctx, data)
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

func (ps *ProjectService) GetProject(ctx context.Context, id string) (res project.Response, err error) {
	data, err := ps.projectRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = project.ParseFromEntity(data)
	return
}

func (ps *ProjectService) DeleteProject(ctx context.Context, id string) (err error) {
	err = ps.projectRepository.Delete(ctx, id)
	return
}

func (ps *ProjectService) UpdateProject(ctx context.Context, id string, req project.Request) (err error) {
	data := project.Entity{
		Title:       req.Title,
		Description: req.Description,
		ManagerID:   project.ParseID(req.ManagerID),
		StartDate:   project.ParseDate(req.StartDate),
		EndDate:     project.ParseDate(req.EndDate),
	}
	err = ps.projectRepository.Update(ctx, id, data)
	return
}

func (ps *ProjectService) SearchProject(ctx context.Context, filter, value string) (res []project.Response, err error) {
	if !project.IsValidFilter(filter) || value == "" {
		err = project.ErrorInvalidSearch
		return
	}
	data, err := ps.projectRepository.Search(ctx, filter, value)
	if err != nil {
		return
	}
	res = project.ParseFromEntities(data)
	return
}

func (ps *ProjectService) GetProjectTasks(ctx context.Context, projectID string) (res []task.Response, err error) {
	_, err = ps.projectRepository.Get(ctx, projectID)
	if err != nil {
		if errors.Is(err, user.ErrorNotFound) {
			return
		}
		return
	}

	data, err := ps.taskRepository.Search(ctx, "project_id", projectID)
	res = task.ParseFromEntities(data)
	if err != nil {
		return
	}
	return
}
