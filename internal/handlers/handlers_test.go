package handlers_test

import (
	"bytes"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/pPrecel/BeerKongServer/internal/auth"
	authMock "github.com/pPrecel/BeerKongServer/internal/auth/mocks"
	composerMock "github.com/pPrecel/BeerKongServer/internal/graphqlcomposer/mocks"
	"github.com/pPrecel/BeerKongServer/internal/handlers"
	"github.com/pPrecel/BeerKongServer/pkg/prisma/generated/prisma-client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_PrismaGQL(t *testing.T) {
	userName := "Fred"

	for testName, testData := range map[string]struct {
		tokenUser   *prisma.User
		userWUI     *prisma.UserWhereUniqueInput
		googleUser  auth.GoogleAccount
		tokenId     string
		tokenIdErr  error
		body        []byte
		expecterErr types.GomegaMatcher
	}{
		"when everything is ok": {
			tokenUser:  &prisma.User{Name: userName},
			userWUI:    &prisma.UserWhereUniqueInput{Name: &userName},
			googleUser: auth.GoogleAccount{Name: userName, Sub: "21312das12"},
			tokenId:    "token",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			//given
			g := gomega.NewWithT(t)
			composer := composerMock.Composer{}
			composer.On("Resolver", &prisma.User{Sub: testData.googleUser.Sub}).Return(nil)
			auth := authMock.Auth{}
			auth.On("GetAccount", testData.tokenId).Return(testData.googleUser, testData.tokenIdErr)

			r, err := http.NewRequest("Post", "/", bytes.NewReader(testData.body))
			if err != nil {
				panic(err)
			}
			r.Header.Add("Authorization", testData.tokenId)
			w := httptest.NewRecorder()

			//when
			handler := handlers.New(&auth, &composer)
			handler.PrismaGQL(w, r)
			resp := w.Result()

			//then
			g.Expect(resp).NotTo(gomega.BeNil())
			g.Expect(resp.Header.Get("Content-Type")).NotTo(gomega.BeNil())
		})
	}
}
