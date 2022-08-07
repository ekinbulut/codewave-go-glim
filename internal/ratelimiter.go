package internal

type RateLimiter struct {
	Bucket *Bucket
}

func NewRateLimiter(bs int) *RateLimiter {
	return &RateLimiter{
		Bucket: NewBucket(bs),
	}
}

func (rl *RateLimiter) GetToken() {
	rl.Bucket.RemoveOne()
}

func (rl *RateLimiter) GetBucketSize() (size int) {
	return rl.Bucket.Size()
}
