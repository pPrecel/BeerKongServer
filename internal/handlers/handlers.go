package handlers

import (
	"fmt"
	h "github.com/99designs/gqlgen/handler"
	"github.com/pPrecel/BeerKongServer/internal/auth"
	"github.com/pPrecel/BeerKongServer/internal/graphqlcomposer"
	"github.com/pPrecel/BeerKongServer/pkg/graphql/generated"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Handler interface {
	PrismaGQL(writer http.ResponseWriter, request *http.Request)
}

type handler struct {
	auth     auth.Auth
	composer graphqlcomposer.Composer
}

func New(auth auth.Auth, composer graphqlcomposer.Composer) Handler {
	return handler{auth: auth, composer: composer}
}

func (s handler) PrismaGQL(writer http.ResponseWriter, request *http.Request) {
	log.Infof("Handle request with authorization")
	token := request.Header.Get("Authorization")
	var user *prisma.User
	if token != "" {
		log.Infof("For token: %s", fmtToken(token))
		res, err := s.auth.GetAccount(token)
		if err != nil {
			log.Info("Unrecognised token")
		} else {
			user = parsUser(&res)
		}
	} else {
		log.Info("Without token")
	}

	h.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: s.composer.Resolver(user)})).ServeHTTP(writer, request)
}

func fmtToken(token string) string {
	if len(token) > 10 {
		return fmt.Sprintf(" For token: \"%s...\"", token[0:10])
	} else {
		return fmt.Sprintf(" For token: \"%s\"", token)
	}
}

func parsUser(account *auth.GoogleAccount) *prisma.User {
	return &prisma.User{
		Sub: account.Sub,
	}
}
