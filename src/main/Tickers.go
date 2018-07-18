// Tickers
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	//Tickers can be stopped like timers.
	//Once a ticker is stopped it won’t receive any more values on its channel.
	//We’ll stop ours after 1600ms.
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("ticker stopped")
}
