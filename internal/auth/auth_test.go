package auth_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/pPrecel/BeerKongServer/internal/auth"
	"github.com/pPrecel/BeerKongServer/internal/auth/mocks"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAuth_GetAccount(t *testing.T) {

	for testName, testData := range map[string]struct {
		tokenId       string
		user          auth.GoogleAccount
		expecterErr   types.GomegaMatcher
		status        int
		responseError error
	}{
		"when is ok": {
			tokenId: "72818231231241231223",
			user: auth.GoogleAccount{
				Email: "wielki@mail.com",
				Name:  "Wielki",
				Sub:   "123",
			},
			expecterErr:   gomega.BeNil(),
			status:        http.StatusOK,
			responseError: nil,
		},
		"when returns error and the forbidden status ": {
			tokenId:       "72818231231241231223",
			user:          auth.GoogleAccount{},
			expecterErr:   gomega.HaveOccurred(),
			status:        http.StatusForbidden,
			responseError: errors.New(""),
		},
		"when returns error": {
			tokenId:       "72818231231241231223",
			user:          auth.GoogleAccount{},
			expecterErr:   gomega.HaveOccurred(),
			status:        http.StatusOK,
			responseError: errors.New(""),
		},
		"when returns the forbidden status": {
			tokenId:       "72818231231241231223",
			user:          auth.GoogleAccount{},
			expecterErr:   gomega.HaveOccurred(),
			status:        http.StatusForbidden,
			responseError: nil,
		},
	} {
		t.Run(testName, func(t *testing.T) {
			// given
			g := gomega.NewWithT(t)
			mockHttpClient := mocks.HttpClient{}
			mockHttpClient.On("Get", fmt.Sprintf("%s%s", auth.GoogleApiURL, testData.tokenId)).Return(googleAccountToResponse(testData.user, testData.status, testData.responseError))

			// when
			auth := auth.New(&mockHttpClient)
			ga, err := auth.GetAccount(testData.tokenId)

			//then
			g.Expect(err).To(testData.expecterErr)
			g.Expect(ga).To(gomega.Equal(testData.user))
		})
	}
}

func googleAccountToResponse(user auth.GoogleAccount, status int, err error) (*http.Response, error) {
	body, _ := json.Marshal(user)
	return &http.Response{
		StatusCode: status,
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
	}, err
}
