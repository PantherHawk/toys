package main

import "fmt"
// find the start and end index between which your target will ie
// move in chunks

func infiniteSearch(nums []int, target int) int {
	s := 0
	e := 1
	
	for nums[e] < target {
		tmp = e + 1
		e = e + (e - s + 1) * 2
		s = tmp
	}
	return binarySearch(nums, target, s, e)
}

func binarySearch(n []int, x, s, e int) int {
	for  s <= e {
		mid := s + (e - s) / 2

		if target < n[mid] {
			e = mid - 1
		} else if target > n[mid] {
			s = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println(infiniteSearch())
}
