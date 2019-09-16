package calcsvc

import (
	"context"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"

	"goa.design/goa/v3/security"

	"calcsvc/account/store"

	"calcsvc/gen/account"
)

var (
	// ErrUnauthorized is the error returned by Login when the request credentials
	// are invalid.
	ErrUnauthorized error = account.Unauthorized("invalid username and password combination")

	// ErrInvalidToken is the error returned when the JWT token is invalid.
	ErrInvalidToken error = account.Unauthorized("invalid token")

	// ErrInvalidTokenScopes is the error returned when the scopes provided in
	// the JWT token claims are invalid.
	ErrInvalidTokenScopes error = account.InvalidScopes("invalid scopes in token")

	// Key is the key used in JWT authentication
	Key = []byte("secret")
)

// JWTAuth implements the authorization logic for service "calc" for the "jwt"
// security scheme.
func (s *calcsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) { return Key, nil })
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, account.InvalidScopes(err.Error())
	}

	store := &accountstore.AccountStore{}
	if exist, err := store.ExistByToken(ctx, token); err != nil {
		s.logger.Printf("Failed to get account because of [%T] %v\n", err, err)
		return ctx, err
	} else {
		if !exist {
			return ctx, fmt.Errorf("Account not found")
		}
	}

	return ctx, nil
}
