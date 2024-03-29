// Code generated by goa v3.0.3, DO NOT EDIT.
//
// calc client HTTP transport
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

// Client lists the calc service endpoint HTTP clients.
type Client struct {
	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Multiply Doer is the HTTP client used to make requests to the multiply
	// endpoint.
	MultiplyDoer goahttp.Doer

	// Devide Doer is the HTTP client used to make requests to the devide endpoint.
	DevideDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the calc service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AddDoer:             doer,
		MultiplyDoer:        doer,
		DevideDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Add returns an endpoint that makes HTTP requests to the calc service add
// server.
func (c *Client) Add() goa.Endpoint {
	var (
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Multiply returns an endpoint that makes HTTP requests to the calc service
// multiply server.
func (c *Client) Multiply() goa.Endpoint {
	var (
		encodeRequest  = EncodeMultiplyRequest(c.encoder)
		decodeResponse = DecodeMultiplyResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildMultiplyRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MultiplyDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "multiply", err)
		}
		return decodeResponse(resp)
	}
}

// Devide returns an endpoint that makes HTTP requests to the calc service
// devide server.
func (c *Client) Devide() goa.Endpoint {
	var (
		encodeRequest  = EncodeDevideRequest(c.encoder)
		decodeResponse = DecodeDevideResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDevideRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DevideDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "devide", err)
		}
		return decodeResponse(resp)
	}
}
