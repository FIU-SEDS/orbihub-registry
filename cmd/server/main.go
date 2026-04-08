package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erielC/orbihub-registry/internal/handler"
	"github.com/erielC/orbihub-registry/internal/middleware"
	"github.com/erielC/orbihub-registry/internal/store"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}
	log.Printf(".env file is found")

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	appsStore := store.NewAppsStore(conn)
	appsHandler := handler.NewAppsHandler(appsStore)

	http.HandleFunc("GET /apps", appsHandler.GetApps)
	http.HandleFunc("GET /apps/{id}", appsHandler.GetAppByID)
	http.HandleFunc("POST /apps", middleware.AuthMiddleware(appsHandler.CreateApp))
	http.HandleFunc("PUT /apps/{id}", middleware.AuthMiddleware(appsHandler.UpdateApp))
	http.HandleFunc("DELETE /apps/{id}", middleware.AuthMiddleware(appsHandler.DeleteApp))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status: Ok!")
	})

	// app_test := model.App{ID: "test", Name: "hello", Description: "test", Version: "1.0.0", Repo: "github", Author: "eriel", Image: "test", CreatedAt: time.Now()}

	log.Println("Starting server at http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
