package main

import (
	"fmt"
)

type Filter map[string]interface{}

func main() {
	// Creating a new filter instance
	filter := make(Filter)

	// Adding key-value pairs to the filter
	filter["name"] = "John Doe"
	filter["age"] = 30
	filter["active"] = true

	// Accessing and using the filter
	name, nameExists := filter["name"].(string)
	age, ageExists := filter["age"].(int)
	active, activeExists := filter["active"].(bool)

	fmt.Println("Name:", name, "Exists:", nameExists)
	fmt.Println("Age:", age, "Exists:", ageExists)
	fmt.Println("Active:", active, "Exists:", activeExists)
}
