package models

type PageData struct {
	Title string
	Files []FileInfo
}

type FileInfo struct {
	Name     string
	IsDir    bool
	FullPath string
}
