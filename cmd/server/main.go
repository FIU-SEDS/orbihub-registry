package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/erielC/orbihub-registry/internal/handler"
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

	log.Println("Starting server at http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
