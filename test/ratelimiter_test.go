package test

import (
	"glim/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenFromBucket(t *testing.T) {

	rl := internal.NewRateLimiter(10)
	for i := 0; i < 3; i++ {
		rl.GetToken()
	}
	var size int = rl.GetBucketSize()
	assert.Equal(t, 7, size)
}

func TestGetTokenThrowFalse(t *testing.T) {
	rl := internal.NewRateLimiter(3)
	var expected bool
	for i := 0; i <= 3; i++ {
		expected = rl.GetToken()
	}
	assert.Equal(t, false, expected)
}
