package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"Software_Development_Assignment/models"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data models.PageData) {
	tmpl = fmt.Sprintf("templates/%s.html", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// isSubdirectory checks if a path is a subdirectory of another.
func isSubdirectory(parent, child string) bool {
	rel, err := filepath.Rel(parent, child)
	if err != nil {
		return false
	}
	return !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}

// isValidFileExtension checks if a file extension is valid.
func isValidFileExtension(filename string) bool {
	allowedExtensions := map[string]bool{
		".txt":  true,
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".pdf":  true,
		// Add more extensions as needed
	}

	// Extract the file extension from the filename
	ext := filepath.Ext(filename)

	// Check if the extracted extension exists in the allowedExtensions map
	return allowedExtensions[ext]
}
