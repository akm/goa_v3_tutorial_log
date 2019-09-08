// Code generated by goa v3.0.3, DO NOT EDIT.
//
// account HTTP server types
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	account "calcsvc/gen/account"
)

// SignupUnauthorizedResponseBody is the type of the "account" service "signup"
// endpoint HTTP response body for the "unauthorized" error.
type SignupUnauthorizedResponseBody string

// NewSignupUnauthorizedResponseBody builds the HTTP response body from the
// result of the "signup" endpoint of the "account" service.
func NewSignupUnauthorizedResponseBody(res account.Unauthorized) SignupUnauthorizedResponseBody {
	body := SignupUnauthorizedResponseBody(res)
	return body
}

// NewSignupPayload builds a account service signup endpoint payload.
func NewSignupPayload() *account.SignupPayload {
	return &account.SignupPayload{}
}