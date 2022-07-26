package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//create test for Bucket
func TestBucket(t *testing.T) {
	b := NewBucket(11)
	assert.Equal(t, 11, b.Size())
}

type Token struct {
}

type Bucket struct {
	Tokens *[]Token
}

func NewBucket(quote int) *Bucket {

	tokens := make([]Token, quote, quote)

	return &Bucket{
		Tokens: &tokens,
	}

}

func (b *Bucket) Size() int {
	return len(*b.Tokens)
}
