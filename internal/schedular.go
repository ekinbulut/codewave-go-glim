package internal

import "time"

type schedular struct {
	second int
	done   chan bool
	ticker *time.Ticker
}

func newSchedular(s int) *schedular {
	return &schedular{
		second: s,
		done:   make(chan bool),
		ticker: nil,
	}
}

func (rl *schedular) Start(fn func()) {

	if rl.second == 0 {
		return
	}

	// confugure ticker for every given second
	rl.ticker = time.NewTicker(time.Duration(rl.second) * time.Second)

	go func() {
		for {
			select {
			case <-rl.done:
				return
			case <-rl.ticker.C:
				fn()
			}
		}
	}()
}

func (rl *schedular) Stop() {

	if rl.ticker != nil {
		rl.ticker.Stop()
		rl.done <- true
	}

}
