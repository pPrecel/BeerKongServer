package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pPrecel/BeerKongServer/internal/graphqlcomposer"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
	"github.com/rs/cors"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/pPrecel/BeerKongServer/internal/auth"
	"github.com/pPrecel/BeerKongServer/internal/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Port           string `envconfig:"default=80"`
	PrismaEndpoint string `envconfig:"default=http://localhost"`
	PrismaSecret   string `envconfig:"default=password"`
}

func main() {
	log.Infof("Start server")

	log.Infof("Read all flags")
	conf, err := readFlags()
	if err != nil {
		log.Fatalf("Flag error: %s", err.Error())
	}

	log.Infof("Set prisma connection for options: endpoint - %s, secret - %s", conf.PrismaEndpoint, conf.PrismaSecret)
	prismaClient := prisma.New(&prisma.Options{
		Endpoint: conf.PrismaEndpoint,
		Secret:   conf.PrismaSecret,
	})

	log.Infof("Create auth communicator")
	authClient := auth.New(&http.Client{})

	log.Infof("Setup chi router")
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
		MaxAge:           300,
	}).Handler)

	log.Infof("Create handlers")
	handlers := handlers.New(authClient, graphqlcomposer.New(prismaClient))

	log.Infof("Set handlers")
	router.HandleFunc("/", handlers.PrismaGQL)
	router.Handle("/playground", handler.Playground("GraphQL playground", "/"))

	log.Infof("Start listening on port \":%s\"", conf.Port)
	if err != http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), router) {
		log.Errorf("Can't listen and serve: \"%s\"", err.Error())
	}
}

func readFlags() (Config, error) {
	cfg := Config{}
	err := envconfig.Init(&cfg)
	return cfg, err
}
