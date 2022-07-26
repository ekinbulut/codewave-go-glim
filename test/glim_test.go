package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"glim/internal"
)

//create test for Bucket
func TestBucketSize(t *testing.T) {
	b := internal.NewBucket(11)
	assert.Equal(t, 11, b.Size())
}

func TestAdd(t *testing.T) {
	b := internal.NewBucket(11)
	b.Add(internal.Token{})
	assert.Equal(t, 1, b.Size())
}
