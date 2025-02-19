package main

import (
	"fmt"
	"sort"
	"sync"
)

func isPrimeNum(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeInRange(ch chan int, start, end int) {
	for i := start; i <= end; i++ {
		if isPrimeNum(i) {
			ch <- i
		}
	}
}

func main() {
	limit := 100
	primeChannel := make(chan int)

	rangeLen := 10

	var primes []int

	go func() {
		for p := range primeChannel {
			//fmt.Println(p)
			primes = append(primes, p)
		}
	}()

	var wg sync.WaitGroup

	for start := 0; start < limit; {
		//end := rangeLen + (start * rangeLen) + 1
		end := start + rangeLen - 1
		if end > limit {
			end = limit
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			primeInRange(primeChannel, start, end)
		}(start, end)

		start = end
	}

	wg.Wait()
	close(primeChannel)

	sort.Ints(primes)

	for _, p := range primes {
		fmt.Println(p)
	}
}
