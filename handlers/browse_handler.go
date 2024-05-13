package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"Software_Development_Assignment/constants"
	"Software_Development_Assignment/models"
)

// BrowseHandler retrieves and displays files in the upload directory.
func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	// Construct the directory path based on the request URL.
	dir := filepath.Join(constants.UploadDirectory, r.URL.Path[1:])
	// Check if the requested directory is a subdirectory of the upload directory.
	if !isSubdirectory(constants.UploadDirectory, dir) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}
	// If the directory is empty, set it to the upload directory.
	if dir == "" {
		dir = constants.UploadDirectory
	}
	// Read the contents of the directory.
	files, err := os.ReadDir(dir)
	if err != nil {
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}

	// Prepare data to render the template.
	data := models.PageData{
		Title: "File Browser",
		Files: make([]models.FileInfo, 0),
	}
	// Iterate through the files and directories, adding them to the data slice.
	for _, file := range files {
		fileInfo := models.FileInfo{
			Name:     file.Name(),
			IsDir:    file.IsDir(),
			FullPath: filepath.Join(r.URL.Path, file.Name()),
		}
		data.Files = append(data.Files, fileInfo)
	}

	// Render the 'browse' template with the data.
	renderTemplate(w, "browse", data)
}
