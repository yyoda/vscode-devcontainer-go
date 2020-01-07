package data

import (
	"github.com/wesovilabs/koazee"
	"github.com/yyoda/vscode-devcontainer-go/internal/model"
	"log"
)

// GetUser .
func GetUser(id int) *model.User {
	users := []model.User{}
	err := db.Select(&users, "SELECT * FROM Users WHERE Id=?", id)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if len(users) > 0 {
		return &users[0]
	}

	return nil
}

// GetAllUser .
func GetAllUser() *[]model.User {
	users := []model.User{}
	err := db.Select(&users, "SELECT * FROM Users")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &users
}

// GetLastUser .
func GetLastUser() *model.User {
	users := GetAllUser()
	user := koazee.StreamOf(*users).Last().Val().(model.User)
	return &user
}
