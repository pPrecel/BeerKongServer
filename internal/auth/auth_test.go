package auth

import (
	"fmt"
	"github.com/pPrecel/BeerKongServer/internal/auth/mocks"
	"net/http"
	"testing"
)

func TestAuth_GetAccount(t *testing.T) {

	for testName, testData := range map[string]struct {
		tokenId string
		user    GoogleAccount
		status  int
	}{
		"should be ok": {
			tokenId: "72818231231241231223",
			user: GoogleAccount{
				Email: "wielki@mail.com",
				Name:  "Wielki",
				Sub:   "123",
			},
			status: http.StatusOK,
		},
	} {
		t.Run(testName, func(t *testing.T) {
			mockHttpClient := mocks.HttpClient{}
			mockHttpClient.On("Get", fmt.Sprintf("%s%s", GoogleApiURL, testData.tokenId)).Return(googleAccountToResponse(testData.user, testData.status))

		})
	}
}

func googleAccountToResponse(user GoogleAccount, status int) (*http.Response, error) {
	//body, _ := json.Marshal(user)
    //io.
	//return &http.Response{
	//	StatusCode: status,
	//	Body:       bytes.NewReader(body).Seek(),
	//}, nil
	return nil, nil
}
