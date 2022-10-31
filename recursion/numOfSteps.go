package main

import "fmt"

func numberOfSteps(n int) int {
	if n == 0 {
		return 0
	}
	if n % 2 == 0 {
		return 1 + numberOfSteps(n / 2)
	}
	return 1 + numberOfSteps(n - 1)
}

func main() {
	fmt.Println(numberOfSteps(8))
}
