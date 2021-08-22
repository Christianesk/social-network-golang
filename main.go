package main

import (
	"log"

	"github.com/Christianesk/social-network-golang/db"
	"github.com/Christianesk/social-network-golang/handlers"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("Unable to connect to db")
		return
	}
	handlers.Handlers()
}
