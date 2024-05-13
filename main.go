package main

import (
	"net/http"

	handlers "Software_Development_Assignment/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.BrowseHandler).Methods("GET")
	r.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("/edit/{file}", handlers.EditHandler).Methods("GET", "POST")
	r.HandleFunc("/rename", handlers.RenameHandler).Methods("POST")
	r.HandleFunc("/delete", handlers.DeleteHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
