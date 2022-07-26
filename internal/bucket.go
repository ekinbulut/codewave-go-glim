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
