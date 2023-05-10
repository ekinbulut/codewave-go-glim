package main

import (
	"fmt"
	"glim/internal"
	"log"
	"time"
)

func main() {

	limiter := internal.NewRateLimiter(5, 250000000, 1)
	config := limiter.Config()
	fmt.Println(config)

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

}
