package internal

import (
	"time"
)

type RateLimiter struct {
	bucket *Bucket
	second int
	done   chan bool
	ticker *time.Ticker
}

func NewRateLimiter(bs int, s int) *RateLimiter {
	return &RateLimiter{
		bucket: NewBucket(bs),
		second: s,
		done:   make(chan bool),
		ticker: nil,
	}
}

func (rl *RateLimiter) GetToken() bool {

	if rl.bucket.Size() > 0 {
		rl.bucket.RemoveOne()
		return true
	}
	return false
}

func (rl *RateLimiter) GetBucketSize() (size int) {
	return rl.bucket.Size()
}

func (rl *RateLimiter) FillBucket() {
	rl.bucket.Fill()
}

func (rl *RateLimiter) Start() {

	if rl.second == 0 {
		return
	}

	// confugure ticker for every given second
	rl.ticker = time.NewTicker(time.Duration(rl.second) * time.Second)

	go func() {
		for {
			select {
			case <-rl.done:
				return
			case <-rl.ticker.C:
				rl.FillBucket()
			}
		}
	}()
}

func (rl *RateLimiter) Stop() {

	if rl.ticker != nil {
		rl.ticker.Stop()
		rl.done <- true
	}

}
