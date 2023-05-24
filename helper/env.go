package helper

import (
	"fmt"
	"os"
)

func GetConnectionString() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Taipei", GetDBHost(), GetDBUser(), GetDBPassword(), GetDBName(), GetDBPort())
}
func GetDBUser() string {
	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		user = "postgres"
	}
	return user
}
func GetDBPassword() string {
	passwd, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		passwd = "postgres"
	}
	return passwd
}
func GetDBName() string {
	name, exists := os.LookupEnv("DB_NAME")
	if !exists {
		name = "postgres"
	}
	return name
}
func GetDBHost() string {
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		host = "localhost"
	}
	return host
}
func GetDBPort() string {
	port, exists := os.LookupEnv("DB_PORT")
	if !exists {
		port = "5432"
	}
	return port
}
