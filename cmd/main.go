package main

import (
	"fmt"
	"github.com/pPrecel/BeerKongServer/internal/auth"
	"github.com/pPrecel/BeerKongServer/internal/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
	"net/http"
)

type Config struct {
	Port string `envconfig:"PORT,default=80"`
}

func main() {
	log.Infof("Start server")

	log.Infof("Read envconfigs")
	var conf Config
	err := envconfig.Init(&conf)
	if err != nil {
		log.Fatalf("Env error: %s", err.Error())
	}

	log.Infof("Create auth communicator")
	authClient := auth.New(&http.Client{})

	log.Infof("Create handlers")
	handlers := handlers.New(authClient)

	log.Infof("Set handlers")
	http.HandleFunc("/", handlers.GraphQlHandler)

	log.Infof("Start listening on port \":%s\"", conf.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)
}
