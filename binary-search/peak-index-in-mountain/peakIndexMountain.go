package main

import "fmt"

func findPeak(nums []int) int {
	max := 0

	e := len(nums) - 1
	s := 0

	for s <= e {
		
		mid := s + (e - s) / 2

		if nums[mid] > max {
			max = nums[mid]
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return max
}

func main() {
	fmt.Println(findPeak([]int{0, 1, 2,3, 20, 1, 0}))
}
