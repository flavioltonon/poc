package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"poc/shared/generic"

	"github.com/labstack/echo/v4"
)

// Notes:
//
// - All handler, middleware and other implementations depend on the contracts predefined by the framework and its echo.Context.
// - Has several boilerplate already implemented, such as JSON encoding/decoding
func main() {
	router := echo.New()

	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			r := c.Request()
			fmt.Printf("got request: %s %s\n", r.Method, r.RequestURI)
			return nil
		})
	})

	router.POST("/foo", func(c echo.Context) error {
		var s generic.Struct

		if err := c.Bind(&s); err != nil {
			return c.JSON(http.StatusBadRequest, []byte(`{"error": "invalid request body"}`))
		}

		return c.JSON(http.StatusNoContent, nil)
	})

	server := httptest.NewServer(router)

	client := server.Client()

	b, _ := json.Marshal(generic.Object)

	response, _ := client.Post(server.URL+"/foo", "application/json", bytes.NewReader(b))

	// response.StatusCode: 204
	fmt.Printf("response.StatusCode: %d\n", response.StatusCode)
}
