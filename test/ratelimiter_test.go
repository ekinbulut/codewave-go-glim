package test

import (
	"glim/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenFromBucket(t *testing.T) {

	rl := NewRateLimiter(10)
	for i := 0; i < 3; i++ {
		rl.GetToken()
	}
	var size int = rl.GetBucketSize()
	assert.Equal(t, 7, size)
}

type RateLimiter struct {
	Bucket *internal.Bucket
}

func NewRateLimiter(bs int) *RateLimiter {
	return &RateLimiter{
		Bucket: internal.NewBucket(bs),
	}
}

func (rl *RateLimiter) GetToken() {
	rl.Bucket.RemoveOne()
}

func (rl *RateLimiter) GetBucketSize() (size int) {
	return rl.Bucket.Size()
}
