package task

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrorNotFound        = errors.New("task not found")
	ErrorInvalidTitle    = errors.New("invalid title")
	ErrorInvalidDate     = errors.New("invalid date format")
	ErrorInvalidUserID   = errors.New("invalid user ID")
	ErrorInvalidProjID   = errors.New("invalid project ID")
	ErrorInvalidPriority = errors.New("invalid priority")
	ErrorInvalidStatus   = errors.New("invalid status")
	ErrorInvalidSearch   = errors.New("invalid search parameters")
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	UserID      string `json:"user_id"`
	ProjectID   string `json:"project_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type Response struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
	UserID      uuid.UUID `json:"user_id"`
	ProjectID   uuid.UUID `json:"project_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Priority:    entity.Priority,
		Status:      entity.Status,
		UserID:      entity.UserID,
		ProjectID:   entity.ProjectID,
		StartDate:   entity.StartDate,
		EndDate:     entity.EndDate,
	}
}

func ParseFromEntities(data []Entity) (res []Response) {
	res = make([]Response, 0)
	for _, entity := range data {
		res = append(res, ParseFromEntity(entity))
	}
	return
}

func (r *Request) Validate() error {
	if r.Title == "" {
		return ErrorInvalidTitle
	}
	if _, err := time.Parse("2006-01-02", r.StartDate); err != nil {
		return ErrorInvalidDate
	}
	if _, err := time.Parse("2006-01-02", r.EndDate); err != nil {
		return ErrorInvalidDate
	}
	if _, err := uuid.Parse(r.UserID); err != nil {
		return ErrorInvalidUserID
	}
	if _, err := uuid.Parse(r.ProjectID); err != nil {
		return ErrorInvalidProjID
	}
	if !isValidPriority(r.Priority) {
		return ErrorInvalidPriority
	}
	if !isValidStatus(r.Status) {
		return ErrorInvalidStatus
	}
	return nil
}

func isValidPriority(priority string) bool {
	validPriorities := map[string]bool{
		"low":    true,
		"medium": true,
		"high":   true,
	}
	return validPriorities[priority]
}

func isValidStatus(status string) bool {
	validStatuses := map[string]bool{
		"active":      true,
		"in_progress": true,
		"completed":   true,
	}
	return validStatuses[status]
}

func IsValidFilter(filter string) bool {
	switch filter {
	case "name", "priority", "status", "user_id", "project_id":
		return true
	default:
		return false
	}
}

func ParseDate(date string) (data time.Time) {
	data, _ = time.Parse("2006-01-02", date)
	return
}
func ParseID(id string) (data uuid.UUID) {
	data, _ = uuid.Parse(id)
	return
}
