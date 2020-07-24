package middlewares

import (
	"context"
	"errors"
	"incidents-api/database"
	"incidents-api/utils"
	"net/http"
)

// APICustomer Type represents the top level API client details.
type APICustomer struct {
	CustomerId   int
	CustomerName string
	APIClientID  string
	APIKey       string
}

// LoadCustomer loads API customer info.
func LoadCustomer(APIClientID string, APIKey string) (*APICustomer, error) {

	c := new(APICustomer)

	err := database.DBCon.QueryRow("SELECT id, customer_name, client_id,api_key  FROM incidents.customers WHERE client_id = $1 AND api_key = $2", APIClientID, APIKey).
		Scan(&c.CustomerId, &c.CustomerName, &c.APIClientID, &c.APIKey)

	if err != nil {
		return nil, err
	}

	return c, nil
}

// BasicAuthentication is a function to Get the HTTP Basic Authentication credentials.
func BasicAuthentication(h http.HandlerFunc) http.HandlerFunc {

	// return the middleware function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the HTTP Basic Authentication credentials.
		APIClientID, APIKey, hasAuth := r.BasicAuth()

		if hasAuth {
			myCustomer, err := LoadCustomer(APIClientID, APIKey)
			if err == nil {

				//Add custom data to context
				ctx := context.WithValue(r.Context(), "customerID", myCustomer.CustomerId)

				// Get new http.Request with the new context
				r = r.WithContext(ctx)

				// execute the actual handler
				h.ServeHTTP(w, r)
			} else {
				// No API customer was found matching Basic Authentication credentials.
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				utils.LogError(r, errors.New("Invalid authentication"))
			}
		} else {
			// No Basic Authentication was provided.
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, "No authentication provided.", http.StatusUnauthorized)
		}

	})
}
