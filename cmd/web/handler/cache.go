package handler

import (
	"fmt"
	"github.com/yyoda/vscode-devcontainer-go/cmd/web/data"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/cache/set", func(w http.ResponseWriter, r *http.Request) {
		data.Set("test", "valuuuuue", time.Duration(10)*time.Second)
		fmt.Fprintf(w, "succeeded")
	})

	http.HandleFunc("/cache/get", func(w http.ResponseWriter, r *http.Request) {
		val := data.Get("test")
		fmt.Fprintf(w, val)
	})
}
