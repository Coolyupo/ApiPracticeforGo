package controller

import (
	"fmt"
	"time"

	"gurubear/helper"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// variable in struct, first letter should be uppercase
type Users struct {
	UserID     uuid.UUID `gorm:"primaryKey;<-:create"`
	Name       string
	CreateTime time.Time
	ModifyTime time.Time
}

func GetUser() []Users {
	db, err := gorm.Open(postgres.Open(helper.GetConnectionString()), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		NoLowerCase: true,
	}})
	if err != nil {
		fmt.Printf("error : %v", err)
	}
	var userA []Users

	// default using lowercase like user
	// no need if change NamingStrategy : NoLowerCase: true
	db.Table("Users").Find(&userA)

	return userA
}

func CreateUser(user Users) bool {
	fmt.Println(user)
	db, err := gorm.Open(postgres.Open(helper.GetConnectionString()), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		NoLowerCase: true,
	}})
	if err != nil {
		fmt.Printf("error : %v", err)
		return false
	}
	user.UserID = uuid.New()
	result := db.Create(&user)
	fmt.Print(result.Error)
	fmt.Print(result.RowsAffected)
	return true
}

func DeleteUser(ID uuid.UUID) bool {
	fmt.Println(ID)

	db, err := gorm.Open(postgres.Open(helper.GetConnectionString()), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		NoLowerCase: true,
	}})
	if err != nil {
		fmt.Printf("error : %v", err)
		return false

	}
	result := db.Delete(&Users{}, ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	return true

}

func UpdateUser(user Users) bool {
	db, err := gorm.Open(postgres.Open(helper.GetConnectionString()), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		NoLowerCase: true,
	}})
	if err != nil {
		fmt.Printf("error : %v", err)
		return false

	}
	lastUser := Users{}
	result := db.Where(Users{UserID: user.UserID}).First(&lastUser)
	if result.Error != nil {
		fmt.Println("should respect exist or not")
		return false
	}
	fmt.Println(user)
	db.Save(&user)
	return true

}
