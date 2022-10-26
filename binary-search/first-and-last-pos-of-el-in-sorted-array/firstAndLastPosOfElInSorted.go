package main

import "fmt"

func getIndices(nums []int, target int) []int {
	// given an array sorted in ascending order
	// return the positions of target integer
	result := []int{-1, -1}
	
	result[0] = search(nums, target, true)
	if result[0] != -1 {
		result[1] = search(nums, target, false)
	}
	return result
}

func search(nums []int, target int, firstIndex bool) int {
	s := 0
	e := len(nums) - 1
	a := -1

	for s <= e {
		mid := s + (e - s) / 2

		if target > nums[mid] {
			s = mid + 1
		} else if target < nums[mid] {
			e = mid - 1
		} else {
			a = mid
			if (firstIndex) {
				e = mid - 1
			} else {
				s = mid + 1
			}
		}
	}
	return a
}

func main() {
	fmt.Println(getIndices([]int{5, 7, 7, 8,8, 10}, 8))
}
