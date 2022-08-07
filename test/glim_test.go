package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"glim/internal"
)

// create test for Bucket
func TestBucketSize(t *testing.T) {
	b := internal.NewBucket(11)
	assert.Equal(t, 11, b.Size())
}

func TestAdd(t *testing.T) {
	b := internal.NewBucket(0)
	b.Add(internal.Token{})
	assert.Equal(t, 1, b.Size())
}

func TestRemoveOne(t *testing.T) {
	b := internal.NewBucket(5)
	b.Add(internal.Token{})
	assert.Equal(t, 6, b.Size())
	b.RemoveOne()
	assert.Equal(t, 5, b.Size())
}
