package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/yyoda/vscode-devcontainer-go/cmd/web/handler"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
