package main

import "net/http"

// Middleware is a type that represents a middleware duh.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// MultipleMiddleware is a helper function that accepts a slice of middleware functions and wraps our original handler
func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {

	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware orders
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped

}
