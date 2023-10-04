package middleware

import (
	"net/http"
	"strings"

	"google.golang.org/api/idtoken"
)

func ValidateGoogleIDToken(validator *idtoken.Validator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")

			idToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

			if _, err := validator.Validate(r.Context(), idToken, ""); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err.Error()))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
