package servererrors

import (
	"encoding/json"
	"net/http"

	"github.com/pPrecel/BeerKongServer/internal/programerrors"
	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func ProgramErrorToResponse(appError programerrors.Error, detailedErrorResponse bool) (status int, body ErrorResponse) {
	httpCode := errorCodeToHttpStatus(appError.Code())
	errorMessage := appError.Error()
	return formatErrorResponse(httpCode, errorMessage, detailedErrorResponse)
}

func errorCodeToHttpStatus(code int) int {
	switch code {
	case programerrors.CodeInternal:
		return http.StatusInternalServerError
	case programerrors.CodeNotFound:
		return http.StatusNotFound
	case programerrors.CodeAlreadyExists:
		return http.StatusConflict
	case programerrors.CodeWrongInput:
		return http.StatusBadRequest
	case programerrors.CodeUpstreamServerCallFailed:
		return http.StatusBadGateway
	case programerrors.CodeAuthenticationFailed:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func formatErrorResponse(httpCode int, errorMessage string, detailedErrorResponse bool) (status int, body ErrorResponse) {
	if isInternalError(httpCode) && !detailedErrorResponse {
		return httpCode, ErrorResponse{httpCode, "Internal error."}
	}
	return httpCode, ErrorResponse{httpCode, errorMessage}
}

func isInternalError(httpCode int) bool {
	return httpCode == http.StatusInternalServerError
}

//SendErrorResponse prepares the http error response and sends it to the client
func SendErrorResponse(apperr programerrors.Error, w http.ResponseWriter) {

	httpcode, resp := ProgramErrorToResponse(apperr, false)

	w.WriteHeader(httpcode)
	respJSON, err := json.Marshal(resp)

	if err != nil {
		marshalerr := programerrors.Internal("Failed to marshal error response: %s \nError body: %s", err, apperr.Error())
		log.Warn(marshalerr)
		return
	}
	w.Write(respJSON)
	return
}
