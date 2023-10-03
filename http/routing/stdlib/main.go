package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"poc/shared/generic"
)

func main() {
	router := generic.HTTPRouter()

	server := httptest.NewServer(router)

	client := server.Client()

	b, _ := json.Marshal(generic.Object)

	response, _ := client.Post(server.URL+generic.HTTPRoute, "application/json", bytes.NewReader(b))

	// response.StatusCode: 204
	fmt.Printf("response.StatusCode: %d\n", response.StatusCode)
}
