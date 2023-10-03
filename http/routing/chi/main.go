package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"poc/shared/generic"

	"github.com/go-chi/chi/v5"
)

// Notes:
//
// - Simple, compatible with Go's standard library definitions (e.g. http.Handler, http.HandlerFunc)
func main() {
	router := chi.NewRouter()
	router.Use(generic.HTTPMiddleware)
	router.Post(generic.HTTPRoute, generic.HTTPHandler)

	server := httptest.NewServer(router)

	client := server.Client()

	b, _ := json.Marshal(generic.Object)

	response, _ := client.Post(server.URL+generic.HTTPRoute, "application/json", bytes.NewReader(b))

	// response.StatusCode: 204
	fmt.Printf("response.StatusCode: %d\n", response.StatusCode)
}
