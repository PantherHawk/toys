package main

import "fmt"

func manyMissingNumbers(nums []int) []int {
	n := len(nums)
	m := []int{}
	i := 0
	// cyclic sort
	for i < n {
		v := nums[i]
		if (i != v - 1 && nums[i] != nums[v - 1]) {
			fmt.Println(v)
			nums[v - 1], nums[i] = nums[i], nums[v - 1]
		} else {
			i++
		}
	}
	for i, v := range(nums) {
		if (i != v - 1) {
			m = append(m, i + 1)
		}
	}
	return m
}

func main() {
	fmt.Println(manyMissingNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
