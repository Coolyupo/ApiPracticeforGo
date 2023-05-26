package repository

import (
	"gurubear/repository/postrepo"
	"gurubear/repository/userrepo"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo *userrepo.UserRepo
	PostRepo *postrepo.PostRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := userrepo.NewUserRepo(db)
	postrepo := postrepo.NewUserRepo(db)
	return &Repositories{
		UserRepo: userRepo,
		PostRepo: postrepo,
	}
}
