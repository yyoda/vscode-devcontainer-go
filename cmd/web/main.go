package main

import (
	"github.com/yyoda/vscode-devcontainer-go/cmd/web/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.SayhelloName) //アクセスのルーティングを設定します。
	err := http.ListenAndServe(":9000", nil)   //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
