package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Christianesk/social-network-golang/middlewares"
	"github.com/Christianesk/social-network-golang/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.VerifyDatabase(routes.RegisterUser)).Methods("POST")
	router.HandleFunc("/login", middlewares.VerifyDatabase(routes.Login)).Methods("POST")
	router.HandleFunc("/profile", middlewares.VerifyDatabase(middlewares.ValidateJWT(routes.Profile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
