package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"project-management/internal/domain/project"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (pr *ProjectRepository) Create(ctx context.Context, data project.Entity) (id string, err error) {
	query := `
		INSERT INTO projects (title, description, manager_id)
		VALUES ($1,$2,$3) RETURNING id;`
	args := []any{data.Title, data.Description, data.ManagerID}
	if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = project.ErrorNotFound
		}
	}
	return
}
func (pr *ProjectRepository) List(ctx context.Context) (projects []project.Entity, err error) {
	query := `SELECT title, description, manager_id, start_date, end_date FROM projects ORDER BY id;`
	err = pr.db.SelectContext(ctx, &projects, query)
	return
}
func (pr *ProjectRepository) Get(ctx context.Context, id string) (data project.Entity, err error) {
	query := `SELECT title, description, manager_id, start_date, end_date FROM projects WHERE id = $1;`

	args := []any{id}

	err = pr.db.GetContext(ctx, &data, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = project.ErrorNotFound
		}
	}

	return
}
func (pr *ProjectRepository) Delete(ctx context.Context, id string) (err error) {
	query := `DELETE FROM projects WHERE id = $1 RETURNING id;`

	args := []any{id}

	if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = project.ErrorNotFound
		}
	}
	return
}
func (pr *ProjectRepository) Update(ctx context.Context, id string, data project.Entity) (err error) {

	return
}
