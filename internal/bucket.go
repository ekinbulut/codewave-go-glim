package internal

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

func (b *Bucket) Add(t Token) {
	*b.Tokens = append(*b.Tokens, t)
}

// remove one token
func (b *Bucket) RemoveOne() {
	*b.Tokens = append((*b.Tokens)[:0], (*b.Tokens)[1:]...)
}
