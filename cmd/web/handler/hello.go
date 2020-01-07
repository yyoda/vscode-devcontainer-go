package handler

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
}
