package task

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID          string    `db:"id" bson:"_id"`
	Title       string    `db:"title" bson:"title"`
	Description string    `db:"description" bson:"description"`
	Priority    string    `db:"priority" bson:"priority"`
	Status      string    `db:"status" bson:"status"`
	UserID      uuid.UUID `db:"user_id" bson:"user_id"`
	ProjectID   uuid.UUID `db:"project_id" bson:"project_id"`
	StartDate   time.Time `db:"start_date" bson:"start_date"`
	EndDate     time.Time `db:"end_date" bson:"end_date"`
}
