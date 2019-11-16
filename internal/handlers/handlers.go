package handlers

import (
	"fmt"
	h "github.com/99designs/gqlgen/handler"
	"github.com/pPrecel/BeerKongServer/internal/auth"
	"github.com/pPrecel/BeerKongServer/internal/programerrors"
	"github.com/pPrecel/BeerKongServer/internal/servererrors"
	"github.com/pPrecel/BeerKongServer/pkg/graphql"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Handler interface {
	GraphQlAuthHandler(http.ResponseWriter,*http.Request)
	GraphQlHandler(http.ResponseWriter, *http.Request)
}

type handler struct {
	auth auth.Auth
}

func New(auth auth.Auth) Handler {
	return handler{auth: auth}
}

func (s handler) GraphQlHandler(writer http.ResponseWriter, request *http.Request) {
	log.Infof("Handle request without authorization")
	h.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}})).ServeHTTP(writer, request)
}

func (s handler) GraphQlAuthHandler(writer http.ResponseWriter, request *http.Request) {

	log.Infof("Handle request with authorization")
	token := request.Header.Get("Authorization")
	if token == "" {
		log.Warnf("Unauthorized connection")
		servererrors.SendErrorResponse(programerrors.AuthenticationFailed("Unauthorized connection"), writer)
		return
	}

	log.Infof(" For token: \"%s...\"", token[0:10])
	res, err := s.auth.GetAccount(token)
	if err != nil {
		log.Warnf("ERROR: \"%s\"", err.Error())
		servererrors.SendErrorResponse(err, writer)
		return
	}

	fmt.Printf("User email: %s for token: \"%s...\"", res.Email, token[0:10])
	h.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}})).ServeHTTP(writer, request)
}

