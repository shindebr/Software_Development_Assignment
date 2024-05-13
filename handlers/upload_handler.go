package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"Software_Development_Assignment/constants"
)

// UploadHandler handles file uploads via HTTP POST requests.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve the uploaded file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file extension
	if !isValidFileExtension(handler.Filename) {
		http.Error(w, "Invalid file extension", http.StatusBadRequest)
		return
	}

	// Read the file contents
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Save the file to the upload directory
	err = os.WriteFile(filepath.Join(constants.UploadDirectory, handler.Filename), fileBytes, 0644)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Redirect to the home page after successful upload
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
