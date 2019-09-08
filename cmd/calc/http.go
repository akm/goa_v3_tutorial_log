package main

import (
	account "calcsvc/gen/account"
	calc "calcsvc/gen/calc"
	accountsvr "calcsvc/gen/http/account/server"
	calcsvr "calcsvc/gen/http/calc/server"
	openapisvr "calcsvc/gen/http/openapi/server"
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, calcEndpoints *calc.Endpoints, accountEndpoints *account.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		calcServer    *calcsvr.Server
		openapiServer *openapisvr.Server
		accountServer *accountsvr.Server
	)
	{
		eh := errorHandler(logger)
		calcServer = calcsvr.New(calcEndpoints, mux, dec, enc, eh)
		openapiServer = openapisvr.New(nil, mux, dec, enc, eh)
		accountServer = accountsvr.New(accountEndpoints, mux, dec, enc, eh)
	}
	// Configure the mux.
	calcsvr.Mount(mux, calcServer)
	openapisvr.Mount(mux)
	accountsvr.Mount(mux, accountServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if debug {
			handler = httpmdlwr.Debug(mux, os.Stdout)(handler)
		}
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	lowMux := http.NewServeMux()
	lowMux.Handle("/static/", http.FileServer(http.Dir("../..")))
	lowMux.Handle("/", handler)

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: lowMux}
	for _, m := range calcServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range openapiServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range accountServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
