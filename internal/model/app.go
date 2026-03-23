package model

import (
	"fmt"
	"net/http"
)

type App struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Repo        string `json:"repo"`
	Author      string `json:"author"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at"`
}

var Apps = []App{
	{
		ID:          "telemetry-viewer",
		Name:        "Orbiview",
		Description: "Real-time rocket telemetry monitoring",
		Version:     "1.0.0",
		Repo:        "https://github.com/FIU-SEDS/Orbiview",
		Author:      "Tomas Mejia",
		Image:       "dashboard_logo.png",
		CreatedAt:   "2026",
	},
}

func GetApps() []App {
	return Apps
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
