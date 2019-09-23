// Code generated by goa v3.0.3, DO NOT EDIT.
//
// account client HTTP transport
//
// Command:
// $ goa gen calcsvc/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the account service endpoint HTTP clients.
type Client struct {
	// Signin Doer is the HTTP client used to make requests to the signin endpoint.
	SigninDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the account service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		SigninDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Signin returns an endpoint that makes HTTP requests to the account service
// signin server.
func (c *Client) Signin() goa.Endpoint {
	var (
		encodeRequest  = EncodeSigninRequest(c.encoder)
		decodeResponse = DecodeSigninResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSigninRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SigninDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("account", "signin", err)
		}
		return decodeResponse(resp)
	}
}
