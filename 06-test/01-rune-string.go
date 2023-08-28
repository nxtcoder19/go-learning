package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Test some rune variable
	testRune()
	//
	// // test loop on map, array
	// testMap()
}

// Test some rune variable
func testRune() {
	g := 'Σ' // rune type, an alias for int32, holds a unicode code point.
	x := '🧠'
	y := '¿'

	f := 3.14195 // float64, an IEEE-754 64-bit floating point number.
	c := 3 + 4i
	fmt.Printf(" %v, %v, %v, %v, %v", g, x, y, f, c)

	// String and rune
	str := "Hello, 世界"
	count := utf8.RuneCountInString(str)
	isValid := utf8.ValidString(str)

	fmt.Printf("The string \"%s\" has %d runes.\n ", str, count)
	fmt.Println("isValid", isValid)

	r := '世'
	runLen := utf8.RuneLen(r)
	fmt.Println("rune length is:", runLen)
	fmt.Println("run is valid? :", utf8.ValidRune(r))

	// Encode and decode
	buf := make([]byte, 3)
	utf8.EncodeRune(buf, r)
	fmt.Println(buf)

	buf1 := []byte{228, 184, 150}
	r, size := utf8.DecodeRune(buf1)
	fmt.Printf("%c occupies %d bytes\n", r, size)

	// Hindi text to be encoded and decoded
	hindiText := "नमस्ते"

	// Encode the Hindi text
	encodedText := make([]byte, utf8.UTFMax*len(hindiText))
	encodedCount := utf8.EncodeRune(encodedText, []rune(hindiText)[0])
	for _, r := range []rune(hindiText)[1:] {
		encodedCount += utf8.EncodeRune(encodedText[encodedCount:], r)
	}

	// Print the encoded byte sequence
	fmt.Printf("Encoded Hindi text: %v\n", encodedText[:encodedCount])

	s := 'A'
	rune_str := rune(s)
	fmt.Println(rune_str, s)

}

// test loop on map, array
func testMap() {
	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		// for each pair in the map, print key and value
		fmt.Printf("key=%s, value=%d\n", key, value)
	}

	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("Hello, %s\n", name)
	}
}
