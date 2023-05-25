package models

import (
	"time"

	"github.com/google/uuid"
)

// variable in struct, first letter should be uppercase
type Users struct {
	UserID     uuid.UUID `gorm:"primaryKey;<-:create"`
	Name       string
	CreateTime time.Time
	ModifyTime time.Time
}
