package routes

import (
	"errors"
	"strings"

	"github.com/Christianesk/social-network-golang/db"
	"github.com/Christianesk/social-network-golang/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUser string

func ProcessToken(token string) (*models.Claim, bool, string, error) {

	signature := []byte("Here-your-signature")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tokenResponse, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return signature, nil
	})
	if err == nil {
		_, found, _ := db.ValidateUser(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}

	if !tokenResponse.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}

	return claims, false, string(""), err
}
