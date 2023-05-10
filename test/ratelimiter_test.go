package test

import (
	"glim/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllowThrowFalse(t *testing.T) {
	rl := internal.NewRateLimiter(3, 10, 0)
	var expected bool
	for i := 0; i <= 3; i++ {
		expected = rl.Allow()
	}
	assert.Equal(t, false, expected)
}
