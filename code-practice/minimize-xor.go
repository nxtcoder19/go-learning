package main

import (
	"fmt"
)

func getMinimizeXor(num1, num2 int) int {
	// Step 1: Count the number of set bits in num2
	countBits := func(n int) int {
		count := 0
		for n > 0 {
			count += n & 1
			n >>= 1
		}
		return count
	}

	setBitsNum2 := countBits(num2)
	fmt.Println(setBitsNum2)

	// Step 2: Construct x by prioritizing set bits in num1
	x := 0
	for i := 31; i >= 0 && setBitsNum2 > 0; i-- {
		if (num1 & (1 << i)) != 0 {
			x |= (1 << i)
			setBitsNum2--
		}
	}

	// Step 3: Add remaining set bits in the lowest positions
	for i := 0; i < 32 && setBitsNum2 > 0; i++ {
		if (x & (1 << i)) == 0 {
			x |= (1 << i)
			setBitsNum2--
		}
	}

	return x
}

func main() {
	// Example 1
	num1 := 3
	num2 := 5
	fmt.Println(getMinimizeXor(num1, num2)) // Output: 3

	//// Example 2
	//num1 = 1
	//num2 = 12
	//fmt.Println(getMinimizeXor(num1, num2)) // Output: 3
	//
	//num1 = 25
	//num2 = 72
	//fmt.Println(getMinimizeXor(num1, num2))
}

//Given two positive integers num1 and num2, find the positive integer x such that:
//
//x has the same number of set bits as num2, and
//The value x XOR num1 is minimal.
//Note that XOR is the bitwise XOR operation.
//
//Return the integer x. The test cases are generated such that x is uniquely determined.
//
//The number of set bits of an integer is the number of 1's in its binary representation.
//
//
//
//Example 1:
//
//Input: num1 = 3, num2 = 5
//Output: 3
//Explanation:
//The binary representations of num1 and num2 are 0011 and 0101, respectively.
//The integer 3 has the same number of set bits as num2, and the value 3 XOR 3 = 0 is minimal.
//Example 2:
//
//Input: num1 = 1, num2 = 12
//Output: 3
//Explanation:
//The binary representations of num1 and num2 are 0001 and 1100, respectively.
//The integer 3 has the same number of set bits as num2, and the value 3 XOR 1 = 2 is minimal.
