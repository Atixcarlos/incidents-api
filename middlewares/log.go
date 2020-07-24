package middlewares

import (
	"incidents-api/utils"
	"net/http"
)

// LogMiddleware is a middleware that logs an API hit in the server logs
func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		h.ServeHTTP(w, r) // call ServeHTTP on the original handler

	})
}
