package main

import (
	"fmt"
	"github.com/pPrecel/BeerKongServer/internal/database"
	"github.com/pPrecel/BeerKongServer/pkg/graphql"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/pPrecel/BeerKongServer/internal/auth"
	"github.com/pPrecel/BeerKongServer/internal/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Port string `envconfig:"default=80"`
	DataConfig database.DataConfig
}

func main() {
	log.Infof("Start server")

	log.Infof("Read all envs")
	conf, err := readFlags()
	if err != nil {
		log.Fatalf("Env error: %s", err.Error())
	}

	//db := database.New(conf.DataConfig)
	//db.Open()

	log.Infof("Create auth communicator")
	authClient := auth.New(&http.Client{})

	log.Infof("Create resolver")
	resolver := &graphql.Resolver{}

	log.Infof("Create handlers")
	handlers := handlers.New(authClient, resolver)

	log.Infof("Set handlers")
	http.Handle("/", handler.Playground("GraphQL playground", "/q-tmp"))
	http.HandleFunc("/q-tmp", handlers.GraphQlHandler)
	http.HandleFunc("/query", handlers.GraphQlAuthHandler)

	log.Infof("Start listening on port \":%s\"", conf.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)
}

func readFlags() (Config, error) {
	cfg := Config{}
	err := envconfig.Init(&cfg)
	return cfg, err
}
