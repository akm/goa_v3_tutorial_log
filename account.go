package calcapi

import (
	account "calcsvc/gen/account"
	"context"
	"log"
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
	return
}
