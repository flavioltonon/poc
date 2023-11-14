package main

import (
	"context"
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
		_ = limiter.Wait(ctx)
		calculator.Increment()
	}
}
