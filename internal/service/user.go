package service

import (
	"context"
	"project-management/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (us *UserService) CreateUser(ctx context.Context) (err error) {
	return nil
}

func (us *UserService) ListUsers(ctx context.Context) (err error) {
	return nil
}

func (us *UserService) GetUser(ctx context.Context) (err error) {
	return nil
}

func (us *UserService) DeleteUser(ctx context.Context) (err error) {
	return nil
}

func (us *UserService) UpdateUser(ctx context.Context) (err error) {
	return nil
}
