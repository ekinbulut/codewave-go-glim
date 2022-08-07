## Glim - A RateLimiter for Go

## Introduction

Glim is a rate limiter for Go. It is a simple, fast, and efficient rate limiter. Applies algorithms; such as: `LeakyBucket`, `TokenBucket`, and `SlidingWindow`.

Currently it support only `TokenBucket` algorithm.


## Installation

```bash
go get -u github.com/ekinbulut/glim
```


## Usage

```go
package main

import (
    "time"
    "github.com/ekinbulut/glim"
)

func main() {
    // gets two parameters: `capacity` and `rate`
    limiter := glim.NewRateLimiter(10, time.Second)
    limiter.Start()
    for i := 0; i < 10; i++ {
        if limiter.GetToken() {
            // do something
        }
    }
    limiter.Stop()
}
```

## Contributing

Contributions are welcome.
Fork the repository and make a pull request.

## License

MIT License