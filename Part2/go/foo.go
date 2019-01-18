// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func number_server(add_number <-chan int, control <-chan int, number chan<- int) {
	var i = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
		case msg_number := <- add_number:
			i += msg_number
		case todo := <- control:
			if todo == GetNumber {
				number <- i
			}
		}
	}
}

func incrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- 1
	}
	finished <- true
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- -1
	}
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// Try doing the exercise both with and without it!

	add_number := make(chan int)
	finished := make(chan bool)
	control := make(chan int)
	number := make(chan int)

	go number_server(add_number, control, number)
	go incrementing(add_number, finished)
	go decrementing(add_number, finished)

	<- finished
	<- finished

	control <- GetNumber
	Println("The magic number is:", <-number)
	control <- Exit
}
