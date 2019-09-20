package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("calc", func() {
	Title("Calculator Service")
	Description("HTTP service for adding numbers, a goa teaser")
	Server("calc", func() {
		// production must be the first host to work as default on GAE
		Host("production", func() { URI("http://localhost") })
		Host("development", func() { URI("http://localhost:8080") })
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

// Service describes a service
var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers")

	Error("unauthorized", String, "Credentials are invalid")
	Error("invalid-scopes", String, "Token scopes are invalid")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
		Response("invalid-scopes", StatusForbidden)
	})

	// Method describes a service method (endpoint)
	Method("add", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Attribute("a", Int, "Left operand")
			Attribute("b", Int, "Right operand")
			// Both attributes must be provided when invoking "add"
			Required("a", "b")
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

	Method("multiply", func() {
		Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			Scope("api:read") // Enforce presence of "api:read" scope in JWT claims.
		})

		Payload(func() {
			Field(1, "a", Int, func() { Description("Left operand") })
			Field(2, "b", Int, func() { Description("Right operand") })
			TokenField(3, "token", String, func() {
				Description("JWT used for authentication")
			})
			Required("a", "b", "token")
		})
		Result(Int)

		HTTP(func() {
			GET("/multiply/{a}/{b}")
			Response(StatusOK)
		})
	})

	Method("devide", func() {
		Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			Scope("api:read")  // Enforce presence of "api:read" scope in JWT claims.
			Scope("api:write") // and "api:write" scopes in JWT claims.
		})

		Payload(func() {
			Field(1, "a", Int, func() { Description("Left operand") })
			Field(2, "b", Int, func() { Description("Right operand") })
			TokenField(3, "token", String, func() {
				Description("JWT used for authentication")
			})
			Required("a", "b", "token")
		})
		Result(Int)

		HTTP(func() {
			GET("/devide/{a}/{b}")
			Response(StatusOK)
		})
	})

})

var _ = Service("openapi", func() {
	// Serve the file with relative path ../../gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json")
})

// BasicAuth defines a security scheme using basic authentication. The scheme
// protects the "signin" action used to create JWTs.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during signin")
	Scope("api:read", "Read-only access")
})

var _ = Service("account", func() {
	Description("Create and delete account")

	Error("unauthorized", String, "Credentials are invalid")
	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	Method("signin", func() {
		Description("Creates a valid JWT")

		// The signin endpoint is secured via basic auth
		Security(BasicAuth)

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			UsernameField(1, "username", String, "Username used to perform signin", func() {
				Example("user")
			})
			PasswordField(2, "password", String, "Password used to perform signin", func() {
				Example("password")
			})
			Required("username", "password")
		})

		Result(String)

		HTTP(func() {
			POST("/signin")
			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})
	})

})
