package main

import (
	"fmt"
)

func main() {
	result := handleSwitchCase(5)
	fmt.Println(result)
}

func handleSwitchCase(Week int) string {
	switch day := Week; day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	default:
		return "weekdays is outside of range"
	}
}
