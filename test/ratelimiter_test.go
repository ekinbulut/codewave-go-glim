package test

import (
	"glim/internal"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenFromBucket(t *testing.T) {
	expected := &internal.Token{}
	var tk internal.Token
	rl := NewRateLimiter(10)
	tk = rl.GetToken()
	assert.True(t, reflect.DeepEqual(tk, expected))
}
