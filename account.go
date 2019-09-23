package calcapi

import (
	account "calcsvc/gen/account"
	"context"
	"log"
	"time"

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

// Creates a valid JWT
func (s *accountsrvc) Signin(ctx context.Context, p *account.SigninPayload) (res string, err error) {
	s.logger.Print("account.signin")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"iat":    time.Now().Unix(),
		"scopes": []string{"api:read"}, // Doesn't include api:write
	})

	s.logger.Printf("user '%s' logged in", p.Username)

	// note that if "SignedString" returns an error then it is returned as
	// an internal error to the client
	t, err := token.SignedString(Key)
	if err != nil {
		return "", err
	}
	return t, nil
}
