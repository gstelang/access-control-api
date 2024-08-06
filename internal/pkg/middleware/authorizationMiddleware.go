package middleware

import (
	"context"
	"net/http"

	"github.com/openfga/go-sdk/client"
)

func AuthorizationMiddleware(fgaClient *client.OpenFgaClient) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		// http.HandlerFunc is an alias for func(w http.ResponseWriter, r *http.Request)
		return func(w http.ResponseWriter, r *http.Request) {
			// TODO:
			// If your doing oAuth check, it can be here.
			// stick your userid after the check
			// ctx := context.WithValue(r.Context(), "userID", userID)
			ctx := context.WithValue(r.Context(), "fgaClient", fgaClient)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
