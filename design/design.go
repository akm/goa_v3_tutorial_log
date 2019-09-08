package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("calc", func() {
	Title("Calculator Service")
	Description("HTTP service for adding numbers, a goa teaser")
	Server("calc", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

// Service describes a service
var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers")
	// Method describes a service method (endpoint)
	Method("add", func() {

		Security(JWTAuth, func() { // Use JWT and an API key to secure this endpoint.
			Scope("api:read")  // Enforce presence of both "api:read"
			Scope("api:write") // and "api:write" scopes in JWT claims.
		})

		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			TokenField(2, "a", Int, func() {
				Description("Left operand")
			})
			TokenField(3, "b", Int, func() {
				Description("Right operand")
			})
			// Both attributes must be provided when invoking "add"
			Required("token", "a", "b")
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(Int)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/add/{a}/{b}")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})
})

var _ = Service("openapi", func() {
	// Serve the file with relative path ../../gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json")
})

var _ = Service("account", func() {
	Description("Create and delete account")

	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	Method("signup", func() {
		Description("Sign up  account with ID token from Google")

		Payload(func() {
			PasswordField(2, "id_token", String, "ID token from google")
			Required("id_token")
		})

		Result(String)

		HTTP(func() {
			POST("/signup")
			Response(StatusOK)
		})

	})

})
