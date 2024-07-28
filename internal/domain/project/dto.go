package project

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrorNotFound         = errors.New("error not found")
	ErrorInvalidSearch    = errors.New("invalid search parameters")
	ErrorInvalidTitle     = errors.New("invalid title")
	ErrorInvalidDate      = errors.New("invalid date format")
	ErrorInvalidManagerID = errors.New("invalid manager ID")
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	ManagerID   string `json:"manager_id"`
}

type Response struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	ManagerID   uuid.UUID `json:"manager_id"`
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		StartDate:   entity.StartDate,
		EndDate:     entity.EndDate,
		ManagerID:   entity.ManagerID,
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
	if r.Description == "" {
		return ErrorInvalidTitle
	}

	if _, err := time.Parse("2006-01-02", r.StartDate); err != nil {
		return ErrorInvalidDate
	}
	if _, err := time.Parse("2006-01-02", r.EndDate); err != nil {
		return ErrorInvalidDate
	}
	if _, err := uuid.Parse(r.ManagerID); err != nil {
		return ErrorInvalidManagerID
	}
	return nil
}

func IsValidFilter(filter string) bool {
	switch filter {
	case "title", "manager_id":
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
