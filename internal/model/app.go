package model

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
