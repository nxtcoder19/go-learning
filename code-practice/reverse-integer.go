package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverseInt(x int) int {
	str := strconv.Itoa(x)
	fmt.Println(str)

	var result string
	if strings.HasPrefix(str, "-") {
		nStr := strings.Replace(str, "-", "", 1)
		fmt.Println(nStr)
		for i := len(nStr) - 1; i >= 0; i-- {
			result += string(nStr[i])
		}
	} else {
		for i := len(str) - 1; i >= 0; i-- {
			result += string(str[i])
		}
	}

	if strings.HasPrefix(result, "0") {
		result = result[1:]
	}

	fmt.Println(result)

	res, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	min := -2147483648
	max := 2147483647

	if x < 0 {
		if -res < min {
			return 0
		} else {
			return -res
		}
	} else {
		if res > max {
			return 0
		} else {
			return res
		}
	}

	//return 0

	//if x < 0 {
	//	return -res
	//}
	//return res

}

func main() {
	a := -123
	fmt.Println(reverseInt(a))
	fmt.Println(reverseInteger(257))
}

func reverseInteger(num int) int {
	reversed := 0

	for num != 0 {
		digit := num % 10              // Extract the last digit
		reversed = reversed*10 + digit // Append the digit to the reversed number
		num /= 10                      // Remove the last digit from num
	}

	return reversed
}
