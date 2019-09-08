package calcapi

import (
	account "calcsvc/gen/account"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/api/oauth2/v2"
)

// account service example implementation.
// The example methods log the requests and return zero values.
type accountsrvc struct {
	logger *log.Logger
}

// NewAccount returns the account service implementation.
func NewAccount(logger *log.Logger) account.Service {
	return &accountsrvc{logger}
}

// Sign up  account with ID token from Google
func (s *accountsrvc) Signup(ctx context.Context, p *account.SignupPayload) (res string, err error) {
	s.logger.Print("account.signup")

	var tokenInfo *oauth2.Tokeninfo
	tokenInfo, err = s.verifyIdToken(p.IDToken)
	if err != nil {
		return
	}

	if b, err := json.MarshalIndent(tokenInfo, "", "  "); err != nil {
		s.logger.Printf("Failed to marshal TokenInfo because of [%T] %v\n", err, err)
		return "", err
	} else {
		s.logger.Printf("TokenInfo: %v\n", string(b))
	}

	res = "dummy-token"
	return
}

func (s *accountsrvc) verifyIdToken(idToken string) (*oauth2.Tokeninfo, error) {
	var httpClient = &http.Client{}
	oauth2Service, err := oauth2.New(httpClient)
	if err != nil {
		s.logger.Printf("Failed to get TokenInfo because of [%T] %v\n", err, err)
		return nil, err
	}
	if tokenInfo, err := oauth2Service.Tokeninfo().IdToken(idToken).Do(); err != nil {
		s.logger.Printf("Failed to get TokenInfo because of [%T] %v\n", err, err)
		return nil, err
	} else {
		return tokenInfo, nil
	}
}
