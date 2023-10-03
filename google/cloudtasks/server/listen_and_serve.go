package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"poc/shared/generic"
	"poc/shared/http/middleware"
	"poc/shared/throughput"

	"github.com/gorilla/mux"
)

func ListenAndServe(showBodyParsingLogs bool) {
	throughputCalculator := throughput.NewCalculator()

	router := mux.NewRouter()

	router.Use(middleware.CalculateThroughput(throughputCalculator))

	router.
		HandleFunc(generic.HTTPRoute, func(w http.ResponseWriter, r *http.Request) {
			var data generic.Struct

			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				log.Fatalf("failed to decode task data: %v\n", err)
				return
			}

			if showBodyParsingLogs {
				log.Printf("[%s] task decoded successfully: %+v", time.Now().Format(time.RFC3339), data)
			}
		}).
		Methods("POST")

	http.ListenAndServe(":8080", router)
}
