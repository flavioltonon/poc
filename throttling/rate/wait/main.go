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

// # Output
//
// ## Limit
//
// 1) Limit: rate.Limit(100 /* RPS */) | Burst: 1
//
// [2023-11-14T18:49:41-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:49:42-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:49:43-03:00] throughput in the last 1s: 99.00 RPS
// [2023-11-14T18:49:44-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:49:45-03:00] throughput in the last 1s: 101.00 RPS
// [2023-11-14T18:49:46-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:49:47-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:49:48-03:00] throughput in the last 1s: 99.00 RPS
// [2023-11-14T18:49:49-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:49:50-03:00] throughput in the last 1s: 100.00 RPS
//
// 2) Limit: rate.Limit(1 /* RPS */) | Burst: 100
//
// [2023-11-14T18:54:26-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:54:27-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:54:28-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:54:29-03:00] throughput in the last 1s: 2.00 RPS
// [2023-11-14T18:54:30-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:54:31-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:54:32-03:00] throughput in the last 1s: 2.00 RPS
// [2023-11-14T18:54:33-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:54:34-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:54:35-03:00] throughput in the last 1s: 1.00 RPS
//
// 3) Limit: rate.Limit(1.0 / 3.0 /* RPS */) | Burst: 1
//
// [2023-11-14T18:56:15-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:56:16-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:56:17-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:56:18-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:56:19-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:56:20-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:56:21-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:56:22-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:56:23-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:56:24-03:00] throughput in the last 1s: 1.00 RPS
//
// ## Every
//
// 1) Limit: rate.Every(1 * time.Second) | Burst: 100
//
// [2023-11-14T18:51:08-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:51:09-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:51:10-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:51:11-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:51:12-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:51:13-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:51:14-03:00] throughput in the last 1s: 2.00 RPS
// [2023-11-14T18:51:15-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:51:16-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:51:17-03:00] throughput in the last 1s: 1.00 RPS
//
// 2) Limit: rate.Every(10 * time.Millisecond) | Burst: 1
//
// [2023-11-14T18:52:37-03:00] throughput in the last 1s: 101.00 RPS
// [2023-11-14T18:52:38-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:52:39-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:52:40-03:00] throughput in the last 1s: 99.00 RPS
// [2023-11-14T18:52:41-03:00] throughput in the last 1s: 101.00 RPS
// [2023-11-14T18:52:42-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:52:43-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:52:44-03:00] throughput in the last 1s: 99.00 RPS
// [2023-11-14T18:52:45-03:00] throughput in the last 1s: 100.00 RPS
// [2023-11-14T18:52:46-03:00] throughput in the last 1s: 101.00 RPS
//
// 2) Limit: rate.Every(3 * time.Second) | Burst: 1
//
// [2023-11-14T18:57:45-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:57:46-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:57:47-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:57:48-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:57:49-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:57:50-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-14T18:57:51-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:57:52-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:57:53-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-14T18:57:54-03:00] throughput in the last 1s: 1.00 RPS
