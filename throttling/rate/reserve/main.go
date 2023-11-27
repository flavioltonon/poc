package main

import (
	"fmt"
	"time"

	"poc/shared/throughput"

	"golang.org/x/time/rate"
)

var (
	// limits represent the rate.Limit that should be applied to time windows. The intention here is
	// to test how rate.Reservation would operate in a hourly rate.Limit change basis.
	//
	// Each key represents an 1-hour window in a 24-hour day, while its value represents the rate limit
	// should be used during that period.
	limits = map[int]rate.Limit{
		8:  4,
		9:  4,
		15: 1,
		16: 1,
		17: 3,
		18: 3,
	}
)

func main() {
	// Initialize an empty rate.Limiter, with zero rate.Limit and burst values. When burst starts with a value
	// greater than 0, the limiter allows <burst> tokens to be reserved.
	limiter := new(rate.Limiter)

	go func() {
		// i ranges from 0 to 24 representing the hours within a day
		for i := 0; i < 24; i++ {
			rateLimit := limits[i]

			fmt.Printf("Setting rate limit as %.2f at %dh\n", rateLimit, i)

			limiter.SetLimit(rateLimit)

			if rateLimit > 0 {
				limiter.SetBurst(1)
			} else {
				limiter.SetBurst(0)
			}

			// In order to simplify tests, each "hour" has a duration of 3 seconds.
			time.Sleep(3 * time.Second)
		}
	}()

	calculator := throughput.NewCalculator()

	// calculator.Calculate outputs the throughput at the last second. It is incremented via calculator.Increment.
	go calculator.Calculate(1 * time.Second /* every */)

	for {
		// Wait until reservation is OK, then sleep for the delay required to get a token.
		// If time.Sleep is not used, the rate limit is not respected.
		for {
			if reservation := limiter.Reserve(); reservation.OK() {
				time.Sleep(reservation.Delay())
				break
			}
		}

		calculator.Increment()
	}
}

// Output:
//
// Setting rate limit as 0.00 at 0h
// [2023-11-27T16:25:00-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:01-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:02-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 1h
// [2023-11-27T16:25:03-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:04-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 2h
// [2023-11-27T16:25:05-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:06-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:07-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 3h
// [2023-11-27T16:25:08-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:09-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:10-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:11-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 4h
// [2023-11-27T16:25:12-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:13-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:14-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 5h
// [2023-11-27T16:25:15-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:16-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:17-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 6h
// [2023-11-27T16:25:18-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:19-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:20-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 7h
// [2023-11-27T16:25:21-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:22-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:23-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 4.00 at 8h
// [2023-11-27T16:25:24-03:00] throughput in the last 1s: 3.00 RPS
// [2023-11-27T16:25:25-03:00] throughput in the last 1s: 4.00 RPS
// [2023-11-27T16:25:26-03:00] throughput in the last 1s: 4.00 RPS
// Setting rate limit as 4.00 at 9h
// [2023-11-27T16:25:27-03:00] throughput in the last 1s: 4.00 RPS
// [2023-11-27T16:25:28-03:00] throughput in the last 1s: 4.00 RPS
// [2023-11-27T16:25:29-03:00] throughput in the last 1s: 4.00 RPS
// Setting rate limit as 0.00 at 10h
// [2023-11-27T16:25:30-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-27T16:25:31-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:32-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 11h
// [2023-11-27T16:25:33-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:34-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:35-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 12h
// [2023-11-27T16:25:36-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:37-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:38-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 13h
// [2023-11-27T16:25:39-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:40-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:41-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 14h
// [2023-11-27T16:25:42-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:43-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:44-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 1.00 at 15h
// [2023-11-27T16:25:45-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:46-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-27T16:25:47-03:00] throughput in the last 1s: 1.00 RPS
// Setting rate limit as 1.00 at 16h
// [2023-11-27T16:25:48-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-27T16:25:49-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-27T16:25:50-03:00] throughput in the last 1s: 1.00 RPS
// Setting rate limit as 3.00 at 17h
// [2023-11-27T16:25:51-03:00] throughput in the last 1s: 1.00 RPS
// [2023-11-27T16:25:52-03:00] throughput in the last 1s: 4.00 RPS
// [2023-11-27T16:25:53-03:00] throughput in the last 1s: 3.00 RPS
// Setting rate limit as 3.00 at 18h
// [2023-11-27T16:25:54-03:00] throughput in the last 1s: 3.00 RPS
// [2023-11-27T16:25:55-03:00] throughput in the last 1s: 3.00 RPS
// [2023-11-27T16:25:56-03:00] throughput in the last 1s: 3.00 RPS
// Setting rate limit as 0.00 at 19h
// [2023-11-27T16:25:57-03:00] throughput in the last 1s: 2.00 RPS
// [2023-11-27T16:25:58-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:25:59-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 20h
// [2023-11-27T16:26:00-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:26:01-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:26:02-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 21h
// [2023-11-27T16:26:03-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:26:04-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:26:05-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 22h
// [2023-11-27T16:26:06-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:26:07-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:26:08-03:00] throughput in the last 1s: 0.00 RPS
// Setting rate limit as 0.00 at 23h
// [2023-11-27T16:26:09-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:26:10-03:00] throughput in the last 1s: 0.00 RPS
// [2023-11-27T16:26:11-03:00] throughput in the last 1s: 0.00 RPS
