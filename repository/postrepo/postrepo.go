package postrepo

import (
	"fmt"
	"gurubear/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (repo *PostRepo) GetPost() []models.Posts {
	var p []models.Posts

	repo.db.Debug().Preload("Users").Find(&p)
	// SELECT * FROM "Users" WHERE "Users"."UserID" = 'cccece27-6839-4b15-995e-142ab96107c9'
	// SELECT * FROM "Posts"

	return p
}

func (repo *PostRepo) CreatePost(post models.Posts) bool {
	post.PostID = uuid.New()
	result := repo.db.Create(&post)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	return true
}
func (repo *PostRepo) UpdatePost(post models.Posts) bool {
	lastPost := models.Posts{}
	result := repo.db.Where(models.Posts{PostID: post.PostID}).First(&lastPost)
	if result.Error != nil {
		fmt.Println("should respect exist or not")
		return false
	}
	repo.db.Save(&post)
	return true
}

func (repo *PostRepo) DeletePost(ID uuid.UUID) bool {
	result := repo.db.Delete(&models.Posts{}, ID)

	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	return true
}
