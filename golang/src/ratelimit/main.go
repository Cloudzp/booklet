package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"math/rand"
	"time"
)

var bus = make(chan int, 1000)

func main() {
	go consumer()
	go producer()
	<-make(chan struct{})
}

func producer() {
	l := rate.NewLimiter(100, 100)
	for {
		i := rand.Int()
		r := l.Reserve()
		if !r.OK() {
			fmt.Println(fmt.Errorf("execeeds limiter's burst ! "))
			continue
		}
		time.Sleep(r.Delay())
		//rage
		bus <- i
	}
}

func consumer() {
	var rate int
	var tick = time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			fmt.Println(fmt.Sprintf("%d/s", rate))
			rate = 0
		default:
			rate++
		}

		<-bus
	}
}
