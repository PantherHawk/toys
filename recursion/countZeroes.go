package main

import "fmt"

func countZeroes(n, c int) int {
	fmt.Println(n)
	if n % 10 == n {
		return c
	}

	if n % 10 == 0 {
		c++
	}
	return countZeroes(n / 10, c)
}

func main() {
	fmt.Println(countZeroes(30210004, 0))
}
