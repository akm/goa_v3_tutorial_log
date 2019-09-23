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

// Creates a valid JWT
func (s *accountsrvc) Signin(ctx context.Context, p *account.SigninPayload) (res string, err error) {
	s.logger.Print("account.signin")
	return
}
