# rate.Limiter.Reserve

## Definitions

- Limit: max frequency of tokens that can be reserved per second.
- Burst: number of tokens that can be reserved at the same second.
- Reservation: a reservation for a token in the future.

## The algorithm

1) Define a map variable `limits` to hold the rate limits that should be applied to every 1-hour-period in a day. Each key
in the variable represents an 1-hour-period while its value indicates the rate limit value (in RPS) that should be used.

2) Initialize an empty rate.Limiter, with zero rate.Limit and burst values. When burst starts with a value
greater than 0, the limiter allows `burst` tokens to be reserved.

3) Spawn a goroutine that emulates the change in hours along a 24-hour-day, and sets the rate limiter's rate value
according to the values defined in the global variable `limits`. In our example, each hour takes only 3 seconds.

4) Create a throughput calculator that prints the throughput in the last second at every second.

5) Enter a loop of requests for `calculator.Increment()` calls limited by the rate limiter.

## Running the example

Change the `limits` variable as desired and run:

> go run main.go

The example logs the current throughput every second.
