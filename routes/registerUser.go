package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Christianesk/social-network-golang/db"
	"github.com/Christianesk/social-network-golang/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error receiving data "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required.", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters long.", 400)
		return
	}

	_, isValidUser, _ := db.ValidateUser(user.Email)

	if isValidUser {
		http.Error(w, "User already exists with the email address: "+user.Email, 400)
		return
	}

	_, status, err := db.InsertRecord(user)

	if err != nil {
		http.Error(w, "An error occurred while trying to register as a user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Record could not be saved.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
