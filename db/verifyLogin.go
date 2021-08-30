package db

import (
	"github.com/Christianesk/social-network-golang/models"
	"golang.org/x/crypto/bcrypt"
)

func VerifyLogin(email string, password string) (models.User, bool) {

	user, found, _ := ValidateUser(email)

	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true

}
