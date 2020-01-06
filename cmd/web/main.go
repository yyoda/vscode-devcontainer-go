package main

import (
	"github.com/yyoda/vscode-devcontainer-go/cmd/web/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Hello)
	http.HandleFunc("/health", handler.Health)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
