package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erielC/orbihub-registry/internal/store"
	"github.com/jackc/pgx/v5"
)

type AppsHandler struct {
	Apps store.AppsStore
}

func NewAppsHandler(appStore store.AppsStore) AppsHandler {
	appHandler := AppsHandler{Apps: appStore}
	return appHandler
}

func (ah AppsHandler) GetApps(w http.ResponseWriter, r *http.Request) {
	apps, err := ah.Apps.GetApps()
	if err != nil {
		log.Println("GetApps error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apps)
}

func (ah AppsHandler) GetAppByID(w http.ResponseWriter, r *http.Request) {
	app, err := ah.Apps.GetAppByID(r.PathValue("id"))
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("Cannot get App ID:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(app)
}
