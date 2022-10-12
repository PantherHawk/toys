package main

import "fmt"

func setMismatch(nums []int) []int {
	i := 0
	n := len(nums)

	for i < n {
		if nums[i] != i + 1 {

			return []int{nums[i], i + 1}
		} else {
			i++
		}
	}
	return []int{}

}

func main() {
	fmt.Println(setMismatch([]int{1, 1}))
}
