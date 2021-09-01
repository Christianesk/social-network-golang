package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Christianesk/social-network-golang/db"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(id)

	if err != nil {
		http.Error(w, "An error occurred while trying to search for the registry."+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
