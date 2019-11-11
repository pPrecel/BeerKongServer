package handlers

import (
	"fmt"
	"github.com/pPrecel/BeerKongServer/internal/auth"
	"github.com/pPrecel/BeerKongServer/internal/programerrors"
	"github.com/pPrecel/BeerKongServer/internal/servererrors"
	"net/http"
)

type Handler interface {
	GraphQlHandler(http.ResponseWriter,*http.Request)
}

type handler struct {
	auth auth.Auth
}

func New(auth auth.Auth) Handler {
	return handler{auth: auth}
}

func (s handler) GraphQlHandler(writer http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("Authorization")

	if token == "" {
		servererrors.SendErrorResponse(programerrors.AuthenticationFailed("Unauthorized connection"), writer)
		return
	}

	res, err := s.auth.GetAccount(token)
	if err != nil {
		servererrors.SendErrorResponse(err, writer)
		return
	}

	fmt.Printf("Email: %s for token: %s", res.Email, token)
	writer.WriteHeader(http.StatusOK)
}
