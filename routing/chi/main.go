package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"poc/shared/generic"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Post("/foo", generic.HTTPHandler)

	server := httptest.NewServer(router)

	client := server.Client()

	b, _ := json.Marshal(generic.Object)

	response, _ := client.Post(server.URL+"/foo", "application/json", bytes.NewReader(b))

	// response.StatusCode: 204
	fmt.Printf("response.StatusCode: %d\n", response.StatusCode)
}
