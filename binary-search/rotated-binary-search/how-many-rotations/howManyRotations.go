package main

import "fmt"

func findPivot(nums []int) int {

	s := 0
	e := len(nums) - 1

	for s < e {
		mid := s + (e - s) / 2
		if mid < e && nums[mid] > nums[mid + 1] {
			return mid
		}
		if mid > s && nums[mid] < nums[mid - 1] {
			return mid - 1
		}
		if nums[mid] == nums[s] && nums[mid] == nums[e] {
			// skip duplicates
			if nums[s] > nums[s + 1] {
				return s
			}
			s++
			
			if nums[e] < nums[e - 1] {
				return e - 1
			}
			e--
		} else if nums[s] < nums[mid] || (nums[s] == nums[mid] && nums[mid] > nums[e]) {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return -1
}

func rotationCount(nums []int) int {
	p := findPivot(nums)

	return p + 1
}

func main() {
	fmt.Println(rotationCount([]int{2, 18, 2, 2, 2, 2}))
}
