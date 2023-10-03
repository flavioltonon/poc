package generic

import "net/http"

var HTTPRouter = func() http.Handler {
	router := http.NewServeMux()
	router.Handle("/route", HTTPMiddleware(http.HandlerFunc(HTTPHandler)))
	return router
}
