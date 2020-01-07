package data

import (
	"github.com/yyoda/vscode-devcontainer-go/cmd/web/model"
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
