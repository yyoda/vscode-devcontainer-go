package handler

import (
	"fmt"
	"net/http"
)

// Hello .
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
