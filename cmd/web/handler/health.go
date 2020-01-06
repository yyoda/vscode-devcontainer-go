package handler

import (
	"fmt"
	"net/http"
)

// Health .
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthy")
}
