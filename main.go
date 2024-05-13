package main

import (
	"net/http"

	handlers "Software_Development_Assignment/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("api/v1/files/", handlers.BrowseHandler).Methods("GET")
	r.HandleFunc("api/v1//upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("api/v1/edit/{file}", handlers.EditHandler).Methods("GET", "POST")
	r.HandleFunc("api/v1/rename", handlers.RenameHandler).Methods("POST")
	r.HandleFunc("api/v1/delete", handlers.DeleteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
