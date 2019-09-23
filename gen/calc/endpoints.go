// Code generated by goa v3.0.3, DO NOT EDIT.
//
// calc endpoints
//
// Command:
// $ goa gen calcsvc/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "calc" service endpoints.
type Endpoints struct {
	Add      goa.Endpoint
	Multiply goa.Endpoint
	Devide   goa.Endpoint
}

// NewEndpoints wraps the methods of the "calc" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Add:      NewAddEndpoint(s),
		Multiply: NewMultiplyEndpoint(s, a.JWTAuth),
		Devide:   NewDevideEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "calc" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Add = m(e.Add)
	e.Multiply = m(e.Multiply)
	e.Devide = m(e.Devide)
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "calc".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddPayload)
		return s.Add(ctx, p)
	}
}

// NewMultiplyEndpoint returns an endpoint function that calls the method
// "multiply" of service "calc".
func NewMultiplyEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MultiplyPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:read"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Multiply(ctx, p)
	}
}

// NewDevideEndpoint returns an endpoint function that calls the method
// "devide" of service "calc".
func NewDevideEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DevidePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:read", "api:write"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Devide(ctx, p)
	}
}
