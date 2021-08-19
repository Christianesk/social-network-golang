package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Christianesk/social-network-golang/middlewares"
	"github.com/Christianesk/social-network-golang/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.VerifyDatabase(routers.RegisterUser)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
