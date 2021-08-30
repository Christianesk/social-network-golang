package jwt

import (
	"time"

	"github.com/Christianesk/social-network-golang/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.User) (string, error) {

	signature := []byte("Here-your-signature")

	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastname":  user.Lastname,
		"birthdate": user.Birthdate,
		"biography": user.Biography,
		"location":  user.Location,
		"website":   user.Website,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(signature)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
