package main

import (
	"glim/internal"
	"log"
	"time"
)

func main() {
	println("Hello, world!")

	limiter := internal.NewRateLimiter(5, 5)

	// call Allow() 3 times
	for i := 0; i < 100; i++ {
		// if Allow() return true, print "allow"
		if limiter.Allow() {
			log.Println("allow")
		} else {
			log.Println("not allow")
			// sleep for 1 second
			time.Sleep(1 * time.Second)
		}
	}

	limiter.Stop()

}
