// Code generated by goa v3.0.3, DO NOT EDIT.
//
// calc HTTP client types
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	calc "calcsvc/gen/calc"
)

// AddInvalidScopesResponseBody is the type of the "calc" service "add"
// endpoint HTTP response body for the "invalid-scopes" error.
type AddInvalidScopesResponseBody string

// AddUnauthorizedResponseBody is the type of the "calc" service "add" endpoint
// HTTP response body for the "unauthorized" error.
type AddUnauthorizedResponseBody string

// MultiplyInvalidScopesResponseBody is the type of the "calc" service
// "multiply" endpoint HTTP response body for the "invalid-scopes" error.
type MultiplyInvalidScopesResponseBody string

// MultiplyUnauthorizedResponseBody is the type of the "calc" service
// "multiply" endpoint HTTP response body for the "unauthorized" error.
type MultiplyUnauthorizedResponseBody string

// DevideInvalidScopesResponseBody is the type of the "calc" service "devide"
// endpoint HTTP response body for the "invalid-scopes" error.
type DevideInvalidScopesResponseBody string

// DevideUnauthorizedResponseBody is the type of the "calc" service "devide"
// endpoint HTTP response body for the "unauthorized" error.
type DevideUnauthorizedResponseBody string

// NewAddInvalidScopes builds a calc service add endpoint invalid-scopes error.
func NewAddInvalidScopes(body AddInvalidScopesResponseBody) calc.InvalidScopes {
	v := calc.InvalidScopes(body)
	return v
}

// NewAddUnauthorized builds a calc service add endpoint unauthorized error.
func NewAddUnauthorized(body AddUnauthorizedResponseBody) calc.Unauthorized {
	v := calc.Unauthorized(body)
	return v
}

// NewMultiplyInvalidScopes builds a calc service multiply endpoint
// invalid-scopes error.
func NewMultiplyInvalidScopes(body MultiplyInvalidScopesResponseBody) calc.InvalidScopes {
	v := calc.InvalidScopes(body)
	return v
}

// NewMultiplyUnauthorized builds a calc service multiply endpoint unauthorized
// error.
func NewMultiplyUnauthorized(body MultiplyUnauthorizedResponseBody) calc.Unauthorized {
	v := calc.Unauthorized(body)
	return v
}

// NewDevideInvalidScopes builds a calc service devide endpoint invalid-scopes
// error.
func NewDevideInvalidScopes(body DevideInvalidScopesResponseBody) calc.InvalidScopes {
	v := calc.InvalidScopes(body)
	return v
}

// NewDevideUnauthorized builds a calc service devide endpoint unauthorized
// error.
func NewDevideUnauthorized(body DevideUnauthorizedResponseBody) calc.Unauthorized {
	v := calc.Unauthorized(body)
	return v
}
