package calcapi

import (
	"context"

	"goa.design/goa/v3/security"

	jwt "github.com/dgrijalva/jwt-go"

	gen "calcsvc/gen/calc"
)

var (
	// ErrUnauthorized is the error returned by Login when the request credentials
	// are invalid.
	ErrUnauthorized error = gen.Unauthorized("invalid username and password combination")

	// ErrInvalidToken is the error returned when the JWT token is invalid.
	ErrInvalidToken error = gen.Unauthorized("invalid token")

	// ErrInvalidTokenScopes is the error returned when the scopes provided in
	// the JWT token claims are invalid.
	ErrInvalidTokenScopes error = gen.InvalidScopes("invalid scopes in token")

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
		return ctx, gen.InvalidScopes(err.Error())
	}
	return ctx, nil
}

// BasicAuth implements the authorization logic for service "account" for the
// "basic" security scheme.
func (s *accountsrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
	if user != "goa" {
		return ctx, ErrUnauthorized
	}
	if pass != "rocks" {
		return ctx, ErrUnauthorized
	}
	return ctx, nil
}
