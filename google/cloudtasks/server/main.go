package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"poc/google/cloudtasks/shared"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

const (
	showDecodeLogs     = false
	showThroughputLogs = true
)

func main() {
	throughputCalculator := newThroughputCalculator()

	router := mux.NewRouter()

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		throughputCalculator.increment()

		if showDecodeLogs {
			var data shared.TaskData

			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				log.Fatalf("failed to decode task data: %v\n", err)
				return
			}

			log.Printf("[%s] task decoded successfully: %+v", time.Now().Format(time.RFC3339), data)
		}
	}).Methods("POST")

	http.ListenAndServe(":8080", router)
}

type throughputCalculator struct {
	counter atomic.Int64
}

func newThroughputCalculator() *throughputCalculator {
	calculator := &throughputCalculator{}

	go func() {
		for range time.Tick(1 * time.Second) {
			if showThroughputLogs {
				fmt.Printf("[%s] current throughput: %d RPS\n", time.Now().Format(time.RFC3339), calculator.current())
			}

			calculator.reset()
		}
	}()

	return calculator
}

func (c *throughputCalculator) current() int64 {
	return c.counter.Load()
}

func (c *throughputCalculator) increment() {
	c.counter.Add(1)
}

func (c *throughputCalculator) reset() {
	c.counter.Store(0)
}
