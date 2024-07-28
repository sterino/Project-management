package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"project-management/internal/domain/project"
	interfaces "project-management/internal/repository/interface"
	"strings"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) interfaces.ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (pr *ProjectRepository) Create(ctx context.Context, data project.Entity) (id string, err error) {
	query := `
		INSERT INTO projects (title, description, start_date, end_date, manager_id)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	args := []any{
		data.Title,
		data.Description,
		data.StartDate,
		data.EndDate,
		data.ManagerID,
	}
	if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = project.ErrorNotFound
		}
	}
	return
}

func (pr *ProjectRepository) List(ctx context.Context) (projects []project.Entity, err error) {
	query := `SELECT * FROM projects ORDER BY id;`
	err = pr.db.SelectContext(ctx, &projects, query)
	return
}

func (pr *ProjectRepository) Get(ctx context.Context, id string) (dest project.Entity, err error) {
	query := `SELECT * FROM projects WHERE id = $1;`
	args := []any{id}
	err = pr.db.GetContext(ctx, &dest, query, args...)
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
	sets, args := pr.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, id)
		query := fmt.Sprintf("UPDATE projects SET %s WHERE id = $%d RETURNING id;", strings.Join(sets, ","), len(args))
		if err = pr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = project.ErrorNotFound
			}
		}
	}
	return
}

func (pr *ProjectRepository) Search(ctx context.Context, filter, value string) (dest []project.Entity, err error) {
	dest = []project.Entity{}
	filter = pr.prepareFilter(filter)
	query := fmt.Sprintf("SELECT * FROM projects WHERE %s = $1", filter)
	err = pr.db.SelectContext(ctx, &dest, query, value)
	if err != nil {
		return
	}
	if len(dest) == 0 {
		err = project.ErrorNotFound
		return
	}
	return
}

func (pr *ProjectRepository) prepareArgs(data project.Entity) (sets []string, args []any) {
	if data.Title != "" {
		args = append(args, data.Title)
		sets = append(sets, fmt.Sprintf("title = $%d", len(args)))
	}
	if data.Description != "" {
		args = append(args, data.Description)
		sets = append(sets, fmt.Sprintf("description = $%d", len(args)))
	}
	if data.ManagerID != uuid.Nil {
		args = append(args, data.ManagerID)
		sets = append(sets, fmt.Sprintf("manager_id = $%d", len(args)))
	}
	if data.StartDate.IsZero() {
		args = append(args, data.StartDate)
		sets = append(sets, fmt.Sprintf("start_date = $%d", len(args)))
	}
	if data.EndDate.IsZero() {
		args = append(args, data.EndDate)
		sets = append(sets, fmt.Sprintf("end_date = $%d", len(args)))
	}
	return
}

func (pr *ProjectRepository) prepareFilter(arg string) string {
	switch arg {
	case "title":
		return "title"
	case "manager_id":
		return "manager_id"
	default:
		return ""
	}
}
