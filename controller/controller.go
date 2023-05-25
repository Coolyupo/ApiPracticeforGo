package controller

import (
	v1 "gurubear/controller/api/v1"
	"gurubear/repository"
)

type Controllers struct {
	UserController *v1.UserController
}

// InitControllers returns a new Controllers
func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		UserController: v1.UserInitController(repositories.UserRepo),
	}
}
