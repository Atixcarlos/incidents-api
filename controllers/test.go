package controllers

import (
	"fmt"
	"net/http"
)

// Hello is an example of simple handler function
func Hello(w http.ResponseWriter, r *http.Request) {
	// Handles top-level page.
	fmt.Fprintf(w, "You are on the home page")
}
