package middlewares

import (
	"net/http"

	"github.com/Christianesk/social-network-golang/routes"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token error !"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
