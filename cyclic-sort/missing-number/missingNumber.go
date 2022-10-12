package main

import "fmt"

// Given a range from [0, n], return the number missing in the range
func missingNumber(nums []int) int {
	n := len(nums)
	i := 0
	// cycle sort
	for i < n {
		v := nums[i]
		//missingSum += v
		if (i != v && v != n) {
			nums[i], nums[v] = nums[v], nums[i]
		}
		i++
	}
	for i, v := range(nums) {
		if (i != v) {
			return i
		}
	}
	return n
}

func main() {
	//fmt.Println([]int{1, 1, 1})
	fmt.Println(missingNumber([]int{0, 2, 1, 3}))
}
