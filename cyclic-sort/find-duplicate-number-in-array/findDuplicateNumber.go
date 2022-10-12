package main

import "fmt"

func findDup(nums []int) int {
	n := len(nums)
	i := 0
	// cyclic sort
	for i < n {
		v := nums[i]
		if (i != v - 1) {
			if (nums[v - 1] == v) {
				return nums[i]
			}
			nums[i], nums[v - 1] = nums[v - 1], nums[i]
		} else {
			i++
		}
	}
	return 0
}

func main() {
	fmt.Println(findDup([]int{3, 1, 3, 4, 2}))
}
