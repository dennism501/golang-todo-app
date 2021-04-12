package todo

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          int
	description string
	slug        string
	complete    bool
	createdAt   time.Time
}
