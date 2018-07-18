// SHA1-Hashes
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	k := "md5 this string"

	h := sha1.New()
	h.Write([]byte(s))

	m := md5.New()
	m.Write([]byte(k))

	//This gets the finalized hash result as a byte slice.
	//The argument to Sum can be used to append to an existing byte slice:
	//it usually isn’t needed.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	fmt.Println(k)
	fmt.Printf("%x\n", m)
}
