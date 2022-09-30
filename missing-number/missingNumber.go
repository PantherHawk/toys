package main

import "fmt"

// Given a range from [0, n], return the number missing in the range
func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	missingSum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	i := 0
	// cycle sort
	for i < n {
		v := nums[i]
		missingSum += v
		i++
	}
	return sum - missingSum

}

func main() {
	//fmt.Println([]int{1, 1, 1})
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 6, 7, 8, 9}))
}
