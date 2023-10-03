package generic

import "net/http"

const HTTPRoute = "/foo"

var HTTPRouter = func() http.Handler {
	router := http.NewServeMux()
	router.Handle(HTTPRoute, HTTPMiddleware(http.HandlerFunc(HTTPHandler)))
	return router
}
