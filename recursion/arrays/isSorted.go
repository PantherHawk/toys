package main

import "fmt"
// given an array of integers, determine if it's sorted in ascending order

func isSorted(a []int, i int) bool {
	fmt.Println(a[i])
	if i == len(a) - 1 {
		return true
	}
	return a[i] < a[i + 1] && isSorted(a, i + 1)
}


func main() {
	fmt.Println("sorted: ", isSorted([]int{1, 2, 3, 4, 5, 6}, 0))
	fmt.Println("unsorted: ", isSorted([]int{1, 3, 5, 2, 8, 7}, 0))
}
