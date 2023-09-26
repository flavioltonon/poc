package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"poc/shared/generic"

	"github.com/gorilla/mux"
)

// Notes:
//
// - Simple, compatible with Go's standard library definitions (e.g. http.Handler, http.HandlerFunc)
func main() {
	router := mux.NewRouter()
	router.Use(generic.HTTPMiddleware)
	router.HandleFunc("/foo", generic.HTTPHandler).Methods("POST")

	server := httptest.NewServer(router)

	client := server.Client()

	b, _ := json.Marshal(generic.Object)

	response, _ := client.Post(server.URL+"/foo", "application/json", bytes.NewReader(b))

	// response.StatusCode: 204
	fmt.Printf("response.StatusCode: %d\n", response.StatusCode)
}
