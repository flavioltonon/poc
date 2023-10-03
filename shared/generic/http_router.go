package generic

import "net/http"

var HTTPRouter = func() http.Handler {
	router := http.NewServeMux()
	router.Handle("/foo", HTTPMiddleware(http.HandlerFunc(HTTPHandler)))
	return router
}
