package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"Software_Development_Assignment/constants"
	"Software_Development_Assignment/models"
)

// EditHandler handles requests to edit files.
func EditHandler(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(constants.UploadDirectory, r.URL.Path[len("/edit/"):])

	// Check if the requested file is within the allowed directory
	if !isSubdirectory(constants.UploadDirectory, filePath) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// If the request method is POST, handle the form submission
	if r.Method == "POST" {
		newContent := r.FormValue(string(content))

		// Write the new content to the file
		err := os.WriteFile(filePath, []byte(newContent), 0644)
		if err != nil {
			http.Error(w, "Error writing to file", http.StatusInternalServerError)
			return
		}
		// Redirect to the homepage after successful file update
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// If the request method is not POST, render the edit template
	data := models.PageData{
		Title: "Edit File",
	}
	renderTemplate(w, "edit", data)
}
