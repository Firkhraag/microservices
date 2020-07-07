package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// getProductIDFromURI gets the id form URL
func getProductIDFromURI(r *http.Request) int {
	// mux handles the error
	// id is always a number
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	return id
}
