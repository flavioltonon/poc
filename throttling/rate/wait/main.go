package main

import (
	"context"
	"log"
	"poc/shared/throughput"
	"time"

	"golang.org/x/time/rate"
)

var (
	limit = rate.Limit(1 /* RPS */)
	burst = 1
)

func main() {
	limiter := rate.NewLimiter(limit, burst)

	calculator := throughput.NewCalculator()
	go calculator.Calculate(1 * time.Second /* every */)

	ctx := context.Background()

	for {
		// Wait returns an error when the context it cancelled or if there are no tokens in the bucket (e.g. when
		// the limiter's limit is set as zero in the runtime).
		if err := limiter.Wait(ctx); err != nil {
			log.Println(err)
			continue
		}

		calculator.Increment()
	}
}
