package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pPrecel/BeerKongServer/internal/database"
	"github.com/pPrecel/BeerKongServer/pkg/graphql"
	"github.com/rs/cors"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/pPrecel/BeerKongServer/internal/auth"
	"github.com/pPrecel/BeerKongServer/internal/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Port string `envconfig:"default=81"`
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

	log.Infof("Setup chi router")
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins: 		[]string{"https://beer-kong.herokuapp.com"},
		AllowedMethods:         []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:         []string{"Authorization"},
		AllowCredentials:       true,
		Debug:                  true,
		MaxAge:					300,
	}).Handler)

	log.Infof("Create handlers")
	handlers := handlers.New(authClient, resolver)

	log.Infof("Set handlers")
	router.Handle("/", handler.Playground("GraphQL playground", "/q-tmp"))
	router.HandleFunc("/q-tmp", handlers.GraphQlHandler)
	router.HandleFunc("/query", handlers.GraphQlAuthHandler)

	log.Infof("Start listening on port \":%s\"", conf.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), router)
	if err != nil {
		log.Errorf("Can't listen and serve: \"%s\"", err.Error())
	}
}

func readFlags() (Config, error) {
	cfg := Config{}
	err := envconfig.Init(&cfg)
	return cfg, err
}
