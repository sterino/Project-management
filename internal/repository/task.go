package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"project-management/internal/domain/task"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (tr *TaskRepository) Create(ctx context.Context, data task.Entity) {

}

func (tr *TaskRepository) Update(ctx context.Context, data task.Entity) {

}

func (tr *TaskRepository) Delete(ctx context.Context, data task.Entity) {

}

func (tr *TaskRepository) Get(ctx context.Context, data task.Entity) {}

func (tr *TaskRepository) List(ctx context.Context, data task.Entity) {}

func (tr *TaskRepository) Search(ctx context.Context, data task.Entity) {}
