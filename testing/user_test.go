package testing

import (
	"gurubear/database"
	"gurubear/repository/userrepo"
	"testing"
)

func TestGetUser(t *testing.T) {
	db := database.GetDatabase()
	userrepo := userrepo.NewUserRepo(db)

	userList := userrepo.GetUser()

	if userList == nil {
		t.Error("No User Data")
	}
}
