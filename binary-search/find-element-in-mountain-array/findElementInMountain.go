package main

// given a bitonic array (like a mountain) of integers, find a target integer
import "fmt"

func getPeak(nums []int) int {
	s := 0
	e := len(nums) - 1

	for s < e {
		mid := s + (e - s) / 2

		if nums[mid] > nums[mid + 1] {
			// we are descending
			e = mid
		} else {
			s = mid + 1
		}
	}
	return s
}

func orderAgnosticBinarySearch(nums []int, x, s, e int) int {
	for s <= e {
		mid := s + (e - s) / 2
		if nums[mid] == x {
			return mid
		}

		if nums[0] < nums[len(nums) - 1] {
			if nums[mid] > x {
				e = mid - 1
			} else {
				s = mid + 1
			}
		} else {
			if nums[mid] > x {
				s = mid + 1
			} else {
				e = mid - 1
			}
		}
	}
	return -1
}

func findElementInMountain(nums []int, target int) int {
	// we're searching so we know we want to find to use binary search
	// find the peak
	peak := getPeak(nums)
	// then take slices from the peak and perform binary searches
	// order agnostic binary search
	ans := orderAgnosticBinarySearch(nums, target, 0, peak)
	// if first pass return -1
	if ans != -1 {
		return ans
	}
	// order agnostic binary search in the descending slice
	return orderAgnosticBinarySearch(nums, target, peak, len(nums) - 1)
}

func main() {
	fmt.Println(findElementInMountain([]int{0, 1, 2, 4, 5, 3, 1}, 3))
}
