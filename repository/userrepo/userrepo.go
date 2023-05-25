package userrepo

import (
	"fmt"
	"gurubear/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) CreateUser(user models.Users) bool {

	user.UserID = uuid.New()
	result := repo.db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	return true
}

func (repo *UserRepo) GetUser() []models.Users {
	var userA []models.Users
	repo.db.Table("Users").Find(&userA)
	return userA
}

func (repo *UserRepo) UpdateUser(user models.Users) bool {
	lastUser := models.Users{}
	result := repo.db.Where(models.Users{UserID: user.UserID}).First(&lastUser)
	if result.Error != nil {
		fmt.Println("should respect exist or not")
		return false
	}
	repo.db.Save(&user)
	return true
}

func (repo *UserRepo) DeleteUser(ID uuid.UUID) bool {

	result := repo.db.Delete(&models.Users{}, ID)

	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	return true
}
