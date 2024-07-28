package project

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID          string    `db:"id" bson:"_id"`
	Title       string    `db:"title" bson:"title"`
	Description string    `db:"description" bson:"description"`
	StartDate   time.Time `db:"start_date" bson:"start_date"`
	EndDate     time.Time `db:"end_date" bson:"end_date"`
	ManagerID   uuid.UUID `db:"manager_id" bson:"manager_id"`
}
