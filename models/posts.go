package models

import (
	"time"

	"github.com/google/uuid"
)

//references, default use ID of Users
//foreignKey, default use UsersID
// since our bad naming convention, we all need to modify it.

type Posts struct {
	PostID     uuid.UUID `gorm:"primaryKey;<-:create"`
	UserID     uuid.UUID
	Users      *Users `gorm:"foreignKey:UserID;references:UserID"`
	Tag        string `gorm:"size:50"`
	Title      string `gorm:"size:50"`
	Content    string
	CreateTime time.Time
	ModifyTime time.Time
}
