package test

import (
	"glim/internal"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenFromBucket(t *testing.T) {

	rl := internal.NewRateLimiter(10, 0)
	for i := 0; i < 3; i++ {
		rl.GetToken()
	}
	var size int = rl.GetBucketSize()
	assert.Equal(t, 7, size)
}

func TestGetTokenThrowFalse(t *testing.T) {
	rl := internal.NewRateLimiter(3, 0)
	var expected bool
	for i := 0; i <= 3; i++ {
		expected = rl.GetToken()
	}
	assert.Equal(t, false, expected)
}

func TestFillBucket(t *testing.T) {
	rl := internal.NewRateLimiter(3, 0)
	var expected bool
	for i := 0; i <= 3; i++ {
		expected = rl.GetToken()
	}

	if !expected {
		rl.FillBucket()
	}

	assert.Equal(t, 3, rl.GetBucketSize())

}

func TestScheduleLimiterToFillBucket(t *testing.T) {

	rl := internal.NewRateLimiter(3, 5)
	rl.Start()
	for i := 0; i <= 3; i++ {
		rl.GetToken()
	}
	wait := time.NewTimer(500 * time.Millisecond)
	<-wait.C

	// sleep for 10 seconds
	time.Sleep(6 * time.Second)
	rl.Stop()

	assert.Equal(t, 3, rl.GetBucketSize())

}
