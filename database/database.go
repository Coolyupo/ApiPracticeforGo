package database

import (
	"fmt"
	"gurubear/pkg/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(helper.GetConnectionString()), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		NoLowerCase: true,
	}})
	if err != nil {
		fmt.Printf("Error open Database : %v", err)
	}
	return db
}
