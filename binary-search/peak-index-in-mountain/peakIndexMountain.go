package main

import "fmt"

func findPeak(nums []int) int {
	//max := 0

	e := len(nums) - 1
	s := 0

	for s < e {
		mid := s + (e - s) / 2

		if nums[mid] > nums[mid + 1] {
			e = mid
		} else {
			s = mid + 1
		}
	}
	return nums[s]
}

func main() {
	fmt.Println(findPeak([]int{1, 2, 3, 5, 6, 4, 3, 2}))
}
