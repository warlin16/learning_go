package chann

import "fmt"

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
