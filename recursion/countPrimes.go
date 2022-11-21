package main

import "fmt"

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	composites := make([]bool, n + 1)
	count := 0
	for i := 2; i * i <= n; i++ {
		// if index is false, i.e. number is prime and therefore not composite, mark all multiples of it as false
		if !composites[i] {
			for j := i * 2; j <= n; j = j + i {
				composites[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !composites[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Printf("there are %v prime numbers less than %v.", countPrimes(6000000000), 6000000000)
}
