package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Files handler
type Files struct {
	log   *log.Logger
	store files.Store
}

// NewFiles creates new files handler with a given logger
func NewFiles(l *log.Logger, s files.Storage) *Files {
	return &Files{l, s}
}

//---------------------------- POST ----------------------------

// AddImage adds image to server
func (fh *Files) AddImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	filename := vars["filename"]

	fh.log("Handle POST id", id, "filename", filename)

	// check if file path is valid
	fh.saveFile(id, fn, w, r)
}
