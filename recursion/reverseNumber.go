package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverseNumber(n int) int {
	if n % 10 == n {
		return n
	}
	return int(math.Pow(10, float64(len(strconv.Itoa(n))-1))) * (n % 10) + reverseNumber(n / 10)
}

func main() {
	fmt.Println(reverseNumber(12345))
}
