package main

import "fmt"

func getIndices(nums []int, target int) []int {
	// given an array sorted in ascending order
	// return the positions of target integer
	s := 0
	e := len(nums) - 1
	
	result := []int{-1, -1}

	for s <= e {

		// find first occurrence of target
		mid := s + (e - s) / 2
		
		if nums[mid] > target {
			e = mid - 1
		} else if nums[mid] < target {
			s = mid + 1
		} else if nums[mid + 1] != target {
			result[0] = mid

			// find second occurrence of target
			e = mid - 1
		} else if nums[mid - 1] != target {
			result[1] = mid
			s = mid + 1
		}
	}
	return result
}

func main() {
	fmt.Println(getIndices([]int{5, 7, 7, 8,8, 10}, 8))
}
