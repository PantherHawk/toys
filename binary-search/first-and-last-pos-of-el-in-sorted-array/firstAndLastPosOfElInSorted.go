package main

import "fmt"

func getIndices(nums []int, target int) []int {
	// given an array sorted in ascending order
	// return the positions of target integer
	s := 0
	e := len(nums) - 1
	
	result := []int{}

	for s < e {
		mid := s + (e -s) / 2

		if target > nums[mid] {
			s = mid + 1
		} else if target < nums[mid] {
			e = mid - 1
		} else {
			above := mid + 1
			below := mid - 1
			result = append(result, mid)
			for target == nums[above] {
				result = append(result, above)
				above++
			}
			for target == nums[below] {
				result = append([]int{below}, result...)
				below--
			}
			return result
		}
	}
	return result
}

func main() {
	fmt.Println(getIndices([]int{5, 7, 7, 8,8, 10}, 8))
}
