package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Christianesk/social-network-golang/db"
	"github.com/Christianesk/social-network-golang/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Incorrect data "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.UpdateRecord(user, IDUser)

	if err != nil {
		http.Error(w, "An error occurred while trying to modify registry. Please try again"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "User's registration has not been modified."+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
