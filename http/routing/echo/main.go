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
// - Has several boilerplate already implemented by echo.Context, such reading JSON from a request body and returning a JSON
// response with an HTTP status code
func main() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			r := c.Request()
			fmt.Printf("got request: %s %s\n", r.Method, r.RequestURI)
			return next(c)
		})
	})

	e.POST(generic.HTTPRoute, func(c echo.Context) error {
		var s generic.Struct

		if err := c.Bind(&s); err != nil {
			return c.JSON(http.StatusBadRequest, []byte(`{"error": "invalid request body"}`))
		}

		return c.JSON(http.StatusNoContent, nil)
	})

	server := httptest.NewServer(e)

	client := server.Client()

	b, _ := json.Marshal(generic.Object)

	response, _ := client.Post(server.URL+generic.HTTPRoute, "application/json", bytes.NewReader(b))

	// response.StatusCode: 204
	fmt.Printf("response.StatusCode: %d\n", response.StatusCode)
}
