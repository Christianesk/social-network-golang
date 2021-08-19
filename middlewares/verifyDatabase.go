package middlewares

import (
	"net/http"

	"github.com/Christianesk/social-network-golang/db"
)

func VerifyDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() {
			http.Error(w, "Database connection error", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
