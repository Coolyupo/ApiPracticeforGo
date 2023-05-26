package controller

import (
	postv1 "gurubear/controller/api/v1/post"
	userv1 "gurubear/controller/api/v1/user"
	"gurubear/repository"
)

type Controllers struct {
	UserController *userv1.Controller
	PostController *postv1.Controller
}

// InitControllers returns a new Controllers
func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		UserController: userv1.InitController(repositories.UserRepo),
		PostController: postv1.InitController(repositories.PostRepo),
	}
}
