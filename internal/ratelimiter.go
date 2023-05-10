package internal

import (
	"fmt"
	"sync"
	"time"
)

type AlgorithmType int

const (
	TokenBucket AlgorithmType = iota
	LeakyBucket
)

type RateLimiter struct {
	capacity       int           // Max capacity of the bucket
	ratePerSecond  int           // Number of requests allowed per second
	tokens         int           // Current number of tokens in the bucket
	lastRefillTime int64         // Unix timestamp of the last refill time
	refillInterval int64         // Number of nanoseconds between refills
	mu             sync.Mutex    // Mutex to protect tokens
	allowMu        sync.Mutex    // Mutex to protect access to the Allow() method
	lastAllowTime  int64         // Unix timestamp of the last allowed request
	algorithm      AlgorithmType // Type of algorithm to use
}

func NewRateLimiter(capacity, ratePerSecond int, algorithm AlgorithmType) *RateLimiter {
	rl := &RateLimiter{
		capacity:       capacity,
		ratePerSecond:  ratePerSecond,
		tokens:         capacity,
		lastRefillTime: time.Now().UnixNano(),
		refillInterval: int64(time.Second) / int64(ratePerSecond),
		algorithm:      algorithm,
	}
	go rl.refill()
	return rl
}

func (rl *RateLimiter) refill() {
	for {
		time.Sleep(time.Duration(rl.refillInterval))
		rl.mu.Lock()
		tokensToAdd := int((time.Now().UnixNano() - rl.lastRefillTime) / rl.refillInterval)
		if tokensToAdd > 0 {
			rl.tokens += tokensToAdd
			if rl.tokens > rl.capacity {
				rl.tokens = rl.capacity
			}
			rl.lastRefillTime = time.Now().UnixNano()
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.allowMu.Lock()
	defer rl.allowMu.Unlock()
	rl.mu.Lock()
	defer rl.mu.Unlock()
	switch rl.algorithm {
	case TokenBucket:
		if rl.tokens > 0 && time.Now().UnixNano()-rl.lastAllowTime >= int64(time.Second)/int64(rl.ratePerSecond) {
			rl.tokens--
			rl.lastAllowTime = time.Now().UnixNano()
			return true
		}
	case LeakyBucket:
		if rl.tokens > 0 {
			rl.tokens--
			rl.lastAllowTime = time.Now().UnixNano()
			return true
		}
	}
	return false
}

func (rl *RateLimiter) Config() string {
	config := fmt.Sprintf("Capacity: %d\nRate per second: %d\nAlgorithm: %s\n", rl.capacity, rl.ratePerSecond, string(rl.algorithm))
	return config
}
