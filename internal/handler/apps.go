package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erielC/orbihub-registry/internal/model"
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

func (ah AppsHandler) CreateApp(w http.ResponseWriter, r *http.Request) {
	var app model.App
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if app.ID == "" {
		http.Error(w, "ID is required to create app", http.StatusBadRequest)
		return
	}
	if err := ah.Apps.CreateApp(app); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (ah AppsHandler) UpdateApp(w http.ResponseWriter, r *http.Request) {
	var app model.App
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updated, err := ah.Apps.UpdateApp(r.PathValue("id"), app)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("UpdateApp error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (ah AppsHandler) DeleteApp(w http.ResponseWriter, r *http.Request) {
	err := ah.Apps.DeleteApp(r.PathValue("id"))
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("DeleteApp error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
