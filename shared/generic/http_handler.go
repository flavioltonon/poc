package generic

import (
	"encoding/json"
	"net/http"
)

var HTTPHandler = func(w http.ResponseWriter, r *http.Request) {
	var s Struct

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid request body"}`))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
