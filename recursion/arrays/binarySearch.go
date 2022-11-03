package main

import "fmt"

func binarySearch(a []int, s, e, t int) int {
	if s > e {
		return -1
	}
	mid := s + (e - s) / 2
	if a[mid] == t {
		return mid
	} else if a[mid] > t {
		return binarySearch(a, s, mid - 1, t)
	} else {
		return binarySearch(a, mid + 1, e, t)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binarySearch(a, 0, len(a) -1, 3))
}
