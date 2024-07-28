package service

import (
	"context"
	"errors"
	"project-management/internal/domain/task"
	"project-management/internal/domain/user"
	interfaces "project-management/internal/repository/interface"
	services "project-management/internal/service/interface"
)

type UserService struct {
	userRepository interfaces.UserRepository
	taskRepository interfaces.TaskRepository
}

func NewUserService(userRepository interfaces.UserRepository, taskRepository interfaces.TaskRepository) services.UserService {
	return &UserService{
		userRepository: userRepository,
		taskRepository: taskRepository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, req user.Request) (id string, err error) {
	data := user.Entity{
		Name:  req.Name,
		Email: req.Email,
		Roles: req.Roles,
	}
	id, err = us.userRepository.Create(ctx, data)
	return
}

func (us *UserService) ListUsers(ctx context.Context) (res []user.Response, err error) {
	data, err := us.userRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	res = user.ParseFromEntities(data)
	return
}

func (us *UserService) GetUser(ctx context.Context, id string) (res user.Response, err error) {
	data, err := us.userRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = user.ParseFromEntity(data)
	return
}

func (us *UserService) DeleteUser(ctx context.Context, id string) (err error) {
	err = us.userRepository.Delete(ctx, id)
	return
}

func (us *UserService) UpdateUser(ctx context.Context, id string, req user.Request) (err error) {
	data := user.Entity{
		Name:  req.Name,
		Email: req.Email,
		Roles: req.Roles,
	}
	err = us.userRepository.Update(ctx, id, data)
	return
}

func (us *UserService) SearchUser(ctx context.Context, filter, value string) (res []user.Response, err error) {
	if !user.IsValidFilter(filter) || value == "" {
		err = user.ErrorInvalidSearch
		return
	}
	data, err := us.userRepository.Search(ctx, filter, value)
	if err != nil {
		return
	}
	res = user.ParseFromEntities(data)
	return
}

func (us *UserService) GetUserTasks(ctx context.Context, userID string) (res []task.Response, err error) {
	_, err = us.userRepository.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, user.ErrorNotFound) {
			return
		}
		return
	}

	data, err := us.taskRepository.Search(ctx, "user_id", userID)
	res = task.ParseFromEntities(data)
	if err != nil {
		return
	}
	return
}
