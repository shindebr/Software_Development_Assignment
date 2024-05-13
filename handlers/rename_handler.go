package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"Software_Development_Assignment/constants"
)

// RenameHandler handles the renaming of files in the upload directory.
func RenameHandler(w http.ResponseWriter, r *http.Request) {
	// Get the old path from the URL query parameter "old".
	oldPath := filepath.Join(constants.UploadDirectory, r.URL.Query().Get("old"))

	// Check if the oldPath is a subdirectory of the UploadDirectory.
	if !isSubdirectory(constants.UploadDirectory, oldPath) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// Get the new name from the URL query parameter "new".
	newName := r.URL.Query().Get("new")

	// Create the new path by joining the UploadDirectory with the new name.
	newPath := filepath.Join(constants.UploadDirectory, newName)

	// Attempt to rename the file from the old path to the new path.
	err := os.Rename(oldPath, newPath)
	if err != nil {
		// If there's an error during renaming, return an internal server error.
		http.Error(w, "Error in renaming file", http.StatusInternalServerError)
		return
	}

	// If renaming is successful, redirect the user to the root ("/").
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
