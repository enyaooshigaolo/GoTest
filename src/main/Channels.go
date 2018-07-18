// Channels
package main

import (
	"fmt"
	"time"
)

//Go by Example: Channel Synchronization
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

//Go by Example: Channel Directions
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	//Go by Example: Channels
	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	msg := <-message

	fmt.Println(msg)

	//Go by Example: Channel Buffering
	messages := make(chan string, 2)

	messages <- "buffer1"
	messages <- "buffer2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	messages2 := make(chan string, 2)

	messages2 <- "buffer1"
	messages2 <- "buffer2"

	buffer1 := <-messages2
	buffer2 := <-messages2

	fmt.Println(buffer1)
	fmt.Println(buffer2)

	//Go by Example: Channel Synchronization
	//This is the function we’ll run in a goroutine.
	//The done channel will be used to notify another goroutine
	//that this function’s work is done.
	done := make(chan bool, 1)
	go worker(done)
	//If you removed the <- done line from this program,
	//the program would exit before the worker even started.
	<-done
	/*	is_done := <-done
		if is_done == true {
			fmt.Println("work is done!")
		} else {
			fmt.Println("work is not done!")
		}
	*/

	//Go by Example: Channel Directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
