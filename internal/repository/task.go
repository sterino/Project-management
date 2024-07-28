package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"project-management/internal/domain/task"
	"strings"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (tr *TaskRepository) Create(ctx context.Context, data task.Entity) (id string, err error) {
	query := `INSERT INTO tasks (title, description, priority, status, user_id, project_id, start_date, end_date) VALUES ($1,$2,$3,$4,$5,$6,$7, $8) RETURNING id;`
	args := []any{
		data.Title,
		data.Description,
		data.Priority,
		data.Status,
		data.UserID,
		data.ProjectID,
		data.StartDate,
		data.EndDate,
	}
	if err = tr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = task.ErrorNotFound
		}
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			return "", task.ErrorNotFound
		}
	}
	return
}

func (tr *TaskRepository) Delete(ctx context.Context, id string) (err error) {
	query := `DELETE FROM tasks WHERE id = $1 RETURNING id;`
	args := []any{id}
	if err = tr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = task.ErrorNotFound
			return
		}
	}

	return
}

func (tr *TaskRepository) Get(ctx context.Context, id string) (dest task.Entity, err error) {
	query := `SELECT * FROM tasks WHERE id = $1;`
	args := []any{id}
	err = tr.db.GetContext(ctx, &dest, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		err = task.ErrorNotFound
	}
	return
}

func (tr *TaskRepository) List(ctx context.Context) (dest []task.Entity, err error) {
	query := `SELECT * FROM tasks ORDER BY id;`
	err = tr.db.SelectContext(ctx, &dest, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = task.ErrorNotFound
		}
	}
	return
}

func (tr *TaskRepository) Search(ctx context.Context, filter string, value string) (dest []task.Entity, err error) {
	dest = []task.Entity{}

	filter = tr.prepareFilter(filter)

	query := fmt.Sprintf("SELECT * FROM tasks WHERE %s = $1;", filter)

	err = tr.db.SelectContext(ctx, &dest, query, value)
	if err != nil {
		return
	}
	if len(dest) == 0 {
		err = task.ErrorNotFound
		return
	}
	return
}

func (tr *TaskRepository) Update(ctx context.Context, id string, data task.Entity) (err error) {
	sets, args := tr.prepareArgs(data)
	if len(args) > 0 {
		args = append(args, id)
		query := fmt.Sprintf("UPDATE tasks SET %s WHERE id = $%d RETURNING id;", strings.Join(sets, ", "), len(args))
		err = tr.db.QueryRowContext(ctx, query, args...).Scan(&id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = task.ErrorNotFound
			}
		}
	}
	return
}

func (tr *TaskRepository) prepareArgs(data task.Entity) (sets []string, args []any) {
	if data.Title != "" {
		sets = append(sets, fmt.Sprintf("title = $%d", len(args)))
		args = append(args, data.Title)
	}
	if data.Description != "" {
		sets = append(sets, fmt.Sprintf("description = $%d", len(args)))
		args = append(args, data.Description)
	}
	if data.Priority != "" {
		sets = append(sets, fmt.Sprintf("priority = $%d", len(args)))
		args = append(args, data.Priority)
	}
	if data.Status != "" {
		sets = append(sets, fmt.Sprintf("status = $%d", len(args)))
		args = append(args, data.Status)
	}
	if data.UserID != uuid.Nil {
		sets = append(sets, fmt.Sprintf("user_id = $%d", len(args)))
		args = append(args, data.UserID)
	}
	if data.ProjectID != uuid.Nil {
		sets = append(sets, fmt.Sprintf("project_id = $%d", len(args)))
		args = append(args, data.ProjectID)
	}
	if data.EndDate.IsZero() {
		sets = append(sets, fmt.Sprintf("end_date = $%d", len(args)))
		args = append(args, data.EndDate)
	}
	return
}

func (tr *TaskRepository) prepareFilter(arg string) string {
	switch arg {
	case "title":
		return "title"
	case "priority":
		return "priority"
	case "status":
		return "status"
	case "user_id":
		return "user_id"
	case "project_id":
		return "project_id"
	}
	return ""
}
