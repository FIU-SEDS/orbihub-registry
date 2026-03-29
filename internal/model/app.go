package model

import (
	"fmt"
	"net/http"
	"time"
)

type App struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Version     string    `json:"version"`
	Repo        string    `json:"repo"`
	Author      string    `json:"author"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
}

func PrintAppInfo(apps []App, w http.ResponseWriter) {
	for _, a := range apps {
		fmt.Fprintf(w, "ID: %s\n", a.ID)
		fmt.Fprintf(w, "Name: %s\n", a.Name)
		fmt.Fprintf(w, "Description: %s\n", a.Description)
		fmt.Fprintf(w, "Version: %s\n", a.Version)
		fmt.Fprintf(w, "Repo: %s\n", a.Repo)
		fmt.Fprintf(w, "Author: %s\n", a.Author)
		fmt.Fprintf(w, "Image: %s\n", a.Image)
		fmt.Fprintf(w, "Created At: %s\n\n", a.CreatedAt)
	}
}
