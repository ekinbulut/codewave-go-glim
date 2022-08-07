package internal

type RateLimiter struct {
	Bucket *Bucket
}

func NewRateLimiter(bs int) *RateLimiter {
	return &RateLimiter{
		Bucket: NewBucket(bs),
	}
}

func (rl *RateLimiter) GetToken() bool {

	if rl.Bucket.Size() > 0 {
		rl.Bucket.RemoveOne()
		return true
	}
	return false
}

func (rl *RateLimiter) GetBucketSize() (size int) {
	return rl.Bucket.Size()
}
