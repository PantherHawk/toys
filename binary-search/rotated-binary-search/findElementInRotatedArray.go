package main

import "fmt"
// find pivot
// you have two sorted arrays, up to pivot and after pivot
// binary search up to pivot
// binary search after pivot if answer not found

func findPivot(nums []int) int {
	// how can we find the pivot in an rotated array?
	// when will you know you have found the pivot?
	// the subsequent index will be less
	// when mid > mid + 1

	s := 0
	e := len(nums) - 1

	for s <= e {
		mid := s + (e - s) / 2

		if mid < e && nums[mid] > nums[mid + 1] {
			return mid
		}
		if mid > s && nums[mid] < nums[mid - 1] {
			return mid - 1
		}
		if nums[s] >= nums[mid] {
			e = mid - 1
		} else if nums[s] < nums[mid] {
			s = mid + 1
		}
	}
	return s
}

func binarySearch(nums []int, x, s, e int) int {
	for s <= e {
		mid := s + (e - s) / 2

		if nums[mid] > x {
			e = mid - 1
		} else if nums[mid] < x {
			s = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func findTargetInRotatedArray(nums []int, x int) int {
	p := findPivot(nums)

	ans := binarySearch(nums, x, 0, p)
	if ans != -1 {
		return ans
	}
	return binarySearch(nums, x, p + 1, len(nums) - 1)
}

func main() {
	fmt.Println(findTargetInRotatedArray([]int{4, 5, 6, 7, 0, 1, 2}, 1))
}
