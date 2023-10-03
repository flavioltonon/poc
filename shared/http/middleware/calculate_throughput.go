package middleware

import (
	"net/http"
	"time"

	"poc/shared/throughput"
)

func CalculateThroughput(calculator *throughput.Calculator) func(next http.Handler) http.Handler {
	go calculator.Calculate(1 * time.Second)

	return func(next http.Handler) http.Handler {
		calculator.Increment()
		return next
	}
}
