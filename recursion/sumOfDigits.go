package main

import "fmt"

func sumOfDigits(n int) int {
	if n % 10 == n{
		return n
	}
	return n % 10 + sumOfDigits(n / 10)
}

func main() {
	fmt.Println(sumOfDigits(1234))
}
