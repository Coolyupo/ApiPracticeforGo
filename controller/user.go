package controller

import (
	"github.com/google/uuid"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetUser() []user {

	userA := []user{{ID: uuid.New().String(), Name: "Alan"}}

	userA = append(userA, user{ID: uuid.New().String(), Name: "Mark"})

	return userA
}
