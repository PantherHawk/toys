package main

// given a bitonic array (like a mountain) of integers, find a target integer
import "fmt"

func findElementInMountain(nums []int, target int) int {
	// we're searching so we know we want to find to use binary search
	peak := -1
	// two binary searches
	// find the peak
	// then take slices from the peak and perform binary searches
	s := 0
	e := len(nums) - 1
	for s < e {
		mid := s + (e - s) / 2

		if nums[mid] > nums[mid + 1] {
			e = mid
		} else {
			s = mid + 1
		}
	}
	peak = s
	ascending := nums[:peak]
	descending := nums[peak:]
	
	s = 0
	e = len(ascending) - 1

	for s <= e {
		mid := s + (e - s) / 2

		if ascending[mid] > target {
			e = mid - 1
		} else if ascending[mid] < target {
			s = mid + 1
		} else {
			return mid
		}
	}

	s = 0
	e = len(descending) - 1

	for s <= e {
		mid := s + (e - s) / 2

		if descending[mid] > target {
			s = mid + 1
		} else if descending[mid] < target {
			e = mid - 1
		} else {
			return len(ascending) + mid
		}
	}
	return -1
}

func main() {
	fmt.Println(findElementInMountain([]int{0, 1, 2, 4, 5, 3, 1}, 3))
}
