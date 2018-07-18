// Epoch
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	p(now)

	millis := nanos / 1000000
	p(secs)
	p(millis)
	p(nanos)
	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))
}
