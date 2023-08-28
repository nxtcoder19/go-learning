package main

import (
	"fmt"
)

func main() {
	// var x [2]int
	// x = [2]int{1, 2};

	// var x = [2]int{1, 2}

	x := [2]int{1, 2}

	fmt.Println(x)

	for i := range x {
		fmt.Printf("type1: index: %d, value: %d\n", i, x[i])
	}

	for i, v := range x {
		fmt.Printf("type2: index: %d, value: %d\n", i, v)
	}

	for _, v := range x {
		fmt.Printf("type3: value: %d\n", v)
	}

	// TODO: slices

	arr := [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	slice := arr[1:4]
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(arr[5:12])
	fmt.Printf("length of slice %d, capacity of slice %v ", len(slice), cap(slice))
	// for _, v := range arr {
	// 	fmt.Printf("value is %d\n", v)
	// }

	// sliceExamples()
	// appendInSlice()
	calculateSliceCap()
}

func sliceExamples() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}

func appendInSlice() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}

func calculateSliceCap() {
	mySlice := make([]int, 3, 5)

	fmt.Println("Length of slice1:", len(mySlice))   // Prints 3
	fmt.Println("Capacity of slice1:", cap(mySlice)) // Prints 5

	// Appending elements to the slice
	for i := 0; i < 8; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(mySlice), cap(mySlice))
	}
}
