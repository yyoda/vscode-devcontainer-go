package handler

import (
	"fmt"
	"github.com/yyoda/vscode-devcontainer-go/internal/data"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := data.GetUser(1)
		fmt.Fprintf(w, strconv.Itoa(user.ID)+":"+user.Name)
	})

	http.HandleFunc("/user/last", func(w http.ResponseWriter, r *http.Request) {
		user := data.GetLastUser()
		fmt.Fprintf(w, strconv.Itoa(user.ID)+":"+user.Name)
	})
}
