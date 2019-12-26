package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

const GoogleApiURL = "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token="

//GoogleAccount contains all information about Google Account
type GoogleAccount struct {
	Alg           string `json:"alg"`
	AtHash        string `json:"at_hash"`
	Aud           string `json:"aud"`
	Azp           string `json:"azp"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Exp           string `json:"exp"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Iat           string `json:"iat"`
	Iss           string `json:"iss"`
	Jti           string `json:"jti"`
	Kid           string `json:"kid"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`
	Typ           string `json:"typ"`
}

//Auth describe functionality of this package
type Auth interface {
	GetAccount(string) (GoogleAccount, error)
}

//HttpClient describe given client used for rest communication
type HttpClint interface {
	Get(string) (*http.Response, error)
}

type auth struct {
	client HttpClint
}

//New - Create and return new auth struct
func New(client HttpClint) Auth {
	return &auth{client: client}
}

//GetAccount - Prepare request to google API and return GoogleAccount struct
func (s *auth) GetAccount(tokenId string) (GoogleAccount, error) {
	res, err := s.client.Get(fmt.Sprintf("%s%s", GoogleApiURL, tokenId))
	if err != nil {
		return GoogleAccount{}, errors.Wrapf(err, "While Sending request")
	}
	if res.StatusCode != http.StatusOK {

		return GoogleAccount{}, fmt.Errorf("Google Call fail with Status %s", res.Status)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	var account GoogleAccount
	err = json.Unmarshal(buf.Bytes(), &account)
	if err != nil {
		return GoogleAccount{}, errors.Wrapf(err, "while parsing response")
	}
	return account, nil
}
