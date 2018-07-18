// Range-over-Channels
//This example also showed that it’s
//possible to close a non-empty channel
//but still have the remaining values be received.
package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}
