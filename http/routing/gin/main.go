package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"poc/shared/generic"

	"github.com/gin-gonic/gin"
)

// Notes:
//
// - Breaks compatibility with Go's standard library http.Handler, but limits itself to defining only a couple of custom types
// gin.Context and gin.HandlerFunc (used both as a handler and middleware)
// - Zero allocation
// - Some controls can only be performed using the global variables in the package. It is not possible to make all configurations
// by creating a custom gin.Engine (e.g. gin.SetMode)
// - Has several boilerplate already implemented by gin.Context, such reading JSON from a request body and returning a JSON
// response with an HTTP status code
func main() {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	engine.Use(func(ctx *gin.Context) {
		r := ctx.Request
		fmt.Printf("got request: %s %s\n", r.Method, r.RequestURI)
	})

	engine.POST("/foo", func(ctx *gin.Context) {
		var s generic.Struct

		if err := ctx.Bind(&s); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.Error{
				Err:  generic.Error,
				Type: gin.ErrorTypePublic,
			})

			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	})

	server := httptest.NewServer(engine)

	client := server.Client()

	b, _ := json.Marshal(generic.Object)

	response, _ := client.Post(server.URL+"/foo", "application/json", bytes.NewReader(b))

	// response.StatusCode: 204
	fmt.Printf("response.StatusCode: %d\n", response.StatusCode)
}
