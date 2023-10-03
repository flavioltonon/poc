package generic

import "net/http"

var HTTPServer = &http.Server{
	Handler: HTTPRouter(),
	Addr:    ":8080",
}
