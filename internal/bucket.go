package internal

type Bucket struct {
	Tokens *[]Token
	quote  int
}

func NewBucket(quote int) *Bucket {

	tokens := make([]Token, quote)

	return &Bucket{
		quote:  quote,
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

func (b *Bucket) Fill() {
	tokens := make([]Token, b.quote)
	for _, t := range tokens {
		b.Add(t)
	}
}
