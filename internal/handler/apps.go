package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erielC/orbihub-registry/internal/store"
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
