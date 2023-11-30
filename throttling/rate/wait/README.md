# rate.Limiter.Wait

## Definitions

- Limit: max frequency of tokens that can be reserved per second.
- Every: reserves a token after every time.Duration
- Burst: number of tokens that can be reserved at the same second.

## Running the example

Change the `limit` and `burst` variables as desired and run:

> go run main.go

The example logs the current throughput every second.
