// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

var i = 0

func number_server(increment, decrement, get chan int) {
	for {
		select {
		case <-increment:
			i++
		case <-decrement:
			i--
		case <-get:
			return
		}

	}

}

func incrementing(inc, done chan int) {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000000; j++ {
		inc <- 0
	}
	done <- 0
}

func decrementing(dec, done chan int) {
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000001; j++ {
		dec <- 0
	}
	done <- 0
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(3)

	inc := make(chan int)
	dec := make(chan int)
	get := make(chan int)
	done := make(chan int)

	// TODO: Spawn both functions as goroutines
	go incrementing(inc, done)
	go decrementing(dec, done)
	go number_server(inc, dec, get)

	<-done
	<-done
	get <- 0

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	//time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
