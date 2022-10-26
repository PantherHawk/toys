package main

import "fmt"

func ceilingofNumber(nums []int, target int) int {
	// binary search
	if target > nums[len(nums) - 1] {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start < end + 1 {
		mid := start + (end - start) / 2
		
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return nums[mid]
		}
	}
	return nums[start]
}

func main() {
	fmt.Println(ceilingofNumber([]int{1, 2, 4, 5, 7, 9}, 8))
}
