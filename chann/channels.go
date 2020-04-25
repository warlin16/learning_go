package chann

import (
	"fmt"
	"time"
)

// TestWithBufferedChannels tests
func TestWithBufferedChannels() {
	// channels are a BLOCKING operation
	// you need a goroutine or to buffer a value when creating a
	// channel to avoid crashing your program
	c := make(chan int, 1)

	c <- 27
	fmt.Println(<-c)
}

// TestChannelsWithGR tests
func TestChannelsWithGR() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println("Value from goroutine", <-c)
}

// The two examples above are of bidirectional channels where you can send and receive
// the exmple below is of a channel that you can only send to
func receivingChann() {
	c := make(chan<- string, 1) // send to channel only
	// c := make(<-chan string) // receiver only channel

	c <- "hello"
	fmt.Println("Channels that only receive ^^")
}

func sendToChann(c chan<- int) {
	c <- 27
}

func receiveFromChann(c <-chan int) {
	fmt.Println(<-c)
}

// SendAndReceive sends to a chann and receives from it without waitgroups or mutexes
func SendAndReceive() {
	start := time.Now()
	c := make(chan int)
	go sendToChann(c)
	receiveFromChann(c)
	// since channels are BLOCKING the function receiveFromChann blocks execution
	// but we started to execute sendToChann in it's own goroutine preventing the app
	// from crashing and needing the assistance from a mutex and or waitgroup
	fmt.Println("Finished!", time.Since(start))
}
