package calcsvc

import (
	account "calcsvc/gen/account"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"calcsvc/account/model"
	"calcsvc/account/store"

	"google.golang.org/api/oauth2/v2"

	jwt "github.com/dgrijalva/jwt-go"
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

var jwtSignatureKey = []byte("goa_v3_tutorial_secret")

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":    time.Date(2019, 9, 1, 12, 0, 0, 0, time.UTC).Unix(),
		"iat":    time.Now().Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	if t, e := token.SignedString(jwtSignatureKey); e != nil {
		err = e
		return
	} else {
		res = t
	}

	m := &accountmodel.Account{
		Email:    tokenInfo.Email,
		JwtToken: res,
	}

	store := &accountstore.AccountStore{}
	if _, e := store.Put(ctx, m); e != nil {
		s.logger.Printf("Failed to save account because of [%T] %v\n", e, e)
		err = e
		return
	}

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
