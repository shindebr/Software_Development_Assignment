package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"Software_Development_Assignment/constants"
)

// DeleteHandler handles HTTP DELETE requests to delete a file.
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Construct the file path from the upload directory and the file name provided in the request.
	filePath := filepath.Join(constants.UploadDirectory, r.URL.Query().Get("file"))

	// Check if the requested file path is a subdirectory of the upload directory.
	if !isSubdirectory(constants.UploadDirectory, filePath) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// Attempt to remove the file from the file system.
	err := os.Remove(filePath)
	if err != nil {
		http.Error(w, "Error deleting file", http.StatusInternalServerError)
		return
	}

	// If the file is successfully deleted, redirect the client to the root URL ("/") with HTTP status code 303 (See Other).
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
