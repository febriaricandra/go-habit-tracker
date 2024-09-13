package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/febriaricandra/go-habit-tracker/config"
	"github.com/febriaricandra/go-habit-tracker/internal/utils"
)

func MethodAllow(allowedMethods ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, method := range allowedMethods {
				if r.Method == method {
					next.ServeHTTP(w, r)
					return
				}
			}
			w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		})
	}
}

func ChainMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

func MiddlewareAuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has a valid JWT token
		token, err := utils.GetBearerToken(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Validate the token
		claims, err := utils.ValidateJWT(token, config.GetSecret())
		fmt.Println(claims)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Add the claims to the request context
		ctx := r.Context()
		type contextKey string

		const claimsKey contextKey = "claims"
		ctx = context.WithValue(ctx, claimsKey, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
