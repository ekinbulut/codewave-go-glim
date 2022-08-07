package test

import (
	"glim/internal"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAllowFromBucket(t *testing.T) {

	rl := internal.NewRateLimiter(10, 0)
	for i := 0; i < 3; i++ {
		rl.Allow()
	}
	var size int = rl.GetBucketSize()
	assert.Equal(t, 7, size)
}

func TestAllowThrowFalse(t *testing.T) {
	rl := internal.NewRateLimiter(3, 0)
	var expected bool
	for i := 0; i <= 3; i++ {
		expected = rl.Allow()
	}
	assert.Equal(t, false, expected)
}

func TestFillBucket(t *testing.T) {
	rl := internal.NewRateLimiter(3, 0)
	var expected bool
	for i := 0; i <= 3; i++ {
		expected = rl.Allow()
	}

	if !expected {
		rl.FillBucket()
	}

	assert.Equal(t, 3, rl.GetBucketSize())

}

func TestScheduleLimiterToFillBucket(t *testing.T) {

	rl := internal.NewRateLimiter(3, 5)

	for i := 0; i <= 3; i++ {
		rl.Allow()
	}
	wait := time.NewTimer(500 * time.Millisecond)
	<-wait.C

	// sleep for 10 seconds
	time.Sleep(6 * time.Second)
	rl.Stop()

	assert.Equal(t, 3, rl.GetBucketSize())

}
