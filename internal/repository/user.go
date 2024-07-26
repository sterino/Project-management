package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"project-management/internal/domain/user"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (tr *UserRepository) Create(ctx context.Context, data user.Entity) {
	return
}

func (tr *UserRepository) Update(ctx context.Context, data user.Entity) {

}

func (tr *UserRepository) Delete(ctx context.Context, data user.Entity) {

}

func (tr *UserRepository) Get(ctx context.Context, data user.Entity) {}

func (tr *UserRepository) List(ctx context.Context, data user.Entity) {}

func (tr *UserRepository) Search(ctx context.Context, data user.Entity) {}
