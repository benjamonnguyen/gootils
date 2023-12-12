package httputil

import (
	"fmt"
	"testing"
	"time"
)

func TestExponentialBackoff(t *testing.T) {
	cfg := ExponentialBackoffConfigs{
		Rate:     3.0,
		Initial:  5 * time.Second,
		Max:      30 * time.Minute,
		Interval: time.Second,
	}
	for i := 0; i < 10; i++ {
		fmt.Println(exponentialBackoff(cfg, i))
	}
}
