package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/erielC/orbihub-registry/internal/model"
)

func main() {
	fmt.Println("Hello World")
	apps := model.GetApps()
	// fmt.Println(apps)

	jsonBytes, err := json.Marshal(apps)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apps)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status: Ok!")
	})

	fmt.Println("Starting server at http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

	fmt.Println()
}
