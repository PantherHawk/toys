package main

import "fmt"

// Given an unsorted int array nums, return the smallest positive integer.

func firstMissingPositive(nums []int) int {
	i := 0
	n := len(nums)

	for i < n {
		correct := nums[i] - 1
		if nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}
	fmt.Println(nums)
	for i, v := range(nums) {
		if v != i + 1 {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}
