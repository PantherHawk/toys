package main

import "fmt"

func floorNumber(nums []int, target int) int {
	s := 0
	e := len(nums) - 1

	for s <= e {
		mid := s + (e - s) / 2

		if target < nums[mid] {
			e = mid - 1
		} else if target > nums[mid] {
			s = mid + 1
		} else {
			return nums[mid]
		}
	}
	return nums[e]
}

func main() {
	fmt.Println(floorNumber([]int{1, 2, 3, 4, 6}, 2))
}
