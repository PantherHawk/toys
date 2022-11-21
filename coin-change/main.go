package main

import "fmt"

// Given an integer array of coins of size n 
// representing different types of currency
// and an integer sum,
// the task is to find the number of ways to make sum
// using different combinations from coins

// e.g.
// INPUT: sum = 4, coins[] = {1, 2, 3}
// OUTPUT: 4
// since {1, 1, 1, 1}, {2, 2}, {2, 1, 1}, {3, 1}
func recurseCoinChange(coins []int, n, sum int) int {
	// base case
	// we have an even distribution of coins and sum has reached zero
	if sum == 0 {
		return 1
	}
	// base case
	// we have an uneven distribution of coins and sum is less than zero
	if sum < 0 {
		return 0
	}
	if n <= 0 {
		return 0
	}
	return recurseCoinChange(coins, n - 1, sum) + recurseCoinChange(coins, n, sum - coins[n - 1])
}
// print all prime numbers up to n
func allPrimes(n int) []int {
	r := []int{}
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			r = append(r, i)
		}
	}
	return r
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(isPrime(4))
	fmt.Println(allPrimes(211))
	//fmt.Println(recurseCoinChange([]int{1, 2, 3}, 3, 4))
}
