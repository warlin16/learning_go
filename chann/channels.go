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

// LoopOverChann ranges over a channel and prints its contents
func LoopOverChann() {
	start := time.Now()
	c := make(chan int)
	go sendToChannel(c)

	for value := range c {
		fmt.Println("All the values in the channel", value)
	}

	fmt.Println("Exiting the function. You ranged over a channel!", time.Since(start))
}

func sendToChannel(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	// IMPORTANT TO REMEMBER YOU MUST ALWAYS CLOSE YOUR CHANNEL
	close(c)
}

// SelectAChann experiment with select stmt
func SelectAChann() {
	start := time.Now()
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("Finishing execution", time.Since(start))
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

// select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of Go.
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel:", v)
		case v := <-o:
			fmt.Println("From the odd channel:", v)
		case v := <-q:
			fmt.Println("From the quit channel", v)
			return
		}
	}
}
