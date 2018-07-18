// Command-Line-Flags
package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numberPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("world:", *wordPtr)
	fmt.Println("numb:", *numberPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar", svar)
	fmt.Println("tail:", flag.Args())
}
