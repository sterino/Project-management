package project

import (
	"errors"
)

var (
	ErrorNotFound = errors.New("error not found")
)

type Entity struct {
	ID          string `db:"id" bson:"_id"`
	Title       string `db:"title" bson:"title"`
	Description string `db:"description" bson:"description"`
	ManagerID   string `db:"manager_id" bson:"manager_id"`
	StartDate   string `db:"start_date" bson:"start_date"`
	EndDate     string `db:"end_date" bson:"end_date"`
}
