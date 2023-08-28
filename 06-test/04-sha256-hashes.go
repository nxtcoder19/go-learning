package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	hashSum := h.Sum(nil)

	hashString := hex.EncodeToString(hashSum)

	// fmt.Println(h)
	// fmt.Println(h.Write([]byte(s)))
	fmt.Println(hashString)
	fmt.Println(s)
	fmt.Printf("%x\n", hashSum)

	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)
}
