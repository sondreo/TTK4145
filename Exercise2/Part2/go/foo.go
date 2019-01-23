// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)

var i int = 0;

func incrementing(chn chan int, finished chan bool) {
	for j := 0; j<1000003; j++ {
		i = <- chn
		i++
		chn <- i
	}
	//TODO: signal that the goroutine is finished
  finished<-true
}

func decrementing(chn chan int, finished chan bool) {
	for j := 0; j<1000000; j++ {
		i = <- chn
		i--
		chn <- i
	}
	//TODO: signal that the goroutine is finished
  finished<-true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels
  chn      := make(chan int, 1)
	finished := make(chan bool, 1)

	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.
	chn <- i

	// TODO: Spawn the required goroutines
  go incrementing(chn, finished)
  go decrementing(chn, finished)

	// TODO: block on finished from both "worker" goroutines
  <- finished 
  <- finished
	Println("The magic number is:", i)
}