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
		} else {
			s = mid + 1
		}
	}
	return -1
}


func findPivotWithDuplicates(nums []int) int {
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
		// if mid == s == e, ignore s and e
		if nums[mid] == nums[s] && nums[mid] == nums[e] {
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
	p := findPivotWithDuplicates(nums)
	
	if p == -1 {
		// if no pivot, array is not rotated, just do binary search
		return binarySearch(nums, x, 0, len(nums) - 1)
	}
	if nums[p] == x {
		return p
	}
	if x > nums[0] {
		return binarySearch(nums, x, 0, p - 1)
	} else {
		return binarySearch(nums, x, p + 1, len(nums) - 1)
	}
}

func main() {
	fmt.Println(findTargetInRotatedArray([]int{2, 2, 2, 3, 9, 1, 2}, 1))
}
