package throughput

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Calculator struct {
	counter atomic.Int64
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Calculate(every time.Duration) {
	for range time.Tick(every) {
		throughput := c.calculate(every)
		fmt.Printf("[%s] throughput in the last %s: %.2f RPS\n", time.Now().Format(time.RFC3339), every, throughput)
		c.reset()
	}
}

func (c *Calculator) calculate(period time.Duration) float64 {
	return float64(c.current()) / period.Seconds()
}

func (c *Calculator) current() int64 {
	return c.counter.Load()
}

func (c *Calculator) Increment() {
	c.counter.Add(1)
}

func (c *Calculator) reset() {
	c.counter.Store(0)
}
