package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"poc/shared/generic"
	"poc/shared/http/middleware"
	"poc/shared/throughput"

	"github.com/gorilla/mux"
	"google.golang.org/api/idtoken"
)

func ListenAndServe(showBodyParsingLogs bool) error {
	ctx := context.Background()

	tokenValidator, err := idtoken.NewValidator(ctx)
	if err != nil {
		return fmt.Errorf("failed to create token validator: %w", err)
	}

	throughputCalculator := throughput.NewCalculator()

	router := mux.NewRouter()

	router.Use(middleware.ValidateGoogleIDToken(tokenValidator))

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

	return http.ListenAndServe(":8080", router)
}
