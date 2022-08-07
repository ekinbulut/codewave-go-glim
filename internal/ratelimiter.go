package internal

type RateLimiter struct {
	bucket    *Bucket
	schedular *schedular
}

func NewRateLimiter(capacity int, rate int) *RateLimiter {
	sch := NewSchedular(rate)
	rt := &RateLimiter{
		bucket:    NewBucket(capacity),
		schedular: sch,
	}
	sch.Start(rt.FillBucket)
	return rt
}

func (rl *RateLimiter) Allow() bool {

	if rl.bucket.Size() > 0 {
		rl.bucket.RemoveOne()
		return true
	}
	return false
}

func (rl *RateLimiter) GetBucketSize() (size int) {
	return rl.bucket.Size()
}

func (rl *RateLimiter) FillBucket() {
	rl.bucket.Fill()
}

func (rl *RateLimiter) Stop() {
	rl.schedular.Stop()
}
