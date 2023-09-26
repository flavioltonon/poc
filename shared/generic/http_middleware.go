package generic

import (
	"fmt"
	"net/http"
)

var HTTPMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got request: %s %s\n", r.Method, r.RequestURI)
	})
}
