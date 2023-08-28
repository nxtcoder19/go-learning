package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

	h, _ := strconv.ParseInt("1A3F", 16, 64)
	fmt.Println(h)

	parseInt()
	parseFloat()
}

func parseInt() {
	intValue, err := strconv.Atoi("12345")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Integer:", intValue)
	}

	// Parse an int64
	int64Value, err := strconv.ParseInt("67890", 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed int64:", int64Value)
	}

	// Parse an unsigned integer (uint)
	uintValue, err := strconv.ParseUint("54321", 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Uint:", uintValue)
	}
}

func parseFloat() {
	float32Value, err := strconv.ParseFloat("3.14159", 32)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed float32:", float32Value)
	}

	// Parse a float64
	float64Value, err := strconv.ParseFloat("2.71828", 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed float64:", float64Value)
	}
}
