package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Christianesk/social-network-golang/db"
	"github.com/Christianesk/social-network-golang/jwt"
	"github.com/Christianesk/social-network-golang/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "User or password are invalid"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	doc, exists := db.VerifyLogin(user.Email, user.Password)

	if exists == false {
		http.Error(w, "User or password are invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(doc)

	if err != nil {
		http.Error(w, "An error occurred while generating the token."+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Setting Cookies
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
