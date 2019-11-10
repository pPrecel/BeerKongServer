package main

import (
	"github.com/pPrecel/BeerKongServer/internal/auth"
	"github.com/pPrecel/BeerKongServer/internal/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
)


func main() {
	authClient := auth.New(&http.Client{})

	handlers := handlers.New(authClient)
	http.HandleFunc("/", handlers.GraphQlHandler)
	http.ListenAndServe(":80", nil)
	log.Info("The server starts listening on port 80")
}
