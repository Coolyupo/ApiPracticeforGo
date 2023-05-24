package controller

import (
	"time"

	"github.com/google/uuid"
)

type user struct {
	UserID     uuid.UUID
	Name       string
	CreateTime time.Time
	ModifyTime time.Time
}

func GetUser() []user {

	userA := []user{{UserID: uuid.New(), Name: "Alan", CreateTime: time.Now(), ModifyTime: time.Now()}}

	userA = append(userA, user{UserID: uuid.New(), Name: "Mark", CreateTime: time.Now(), ModifyTime: time.Now()})

	return userA
}
