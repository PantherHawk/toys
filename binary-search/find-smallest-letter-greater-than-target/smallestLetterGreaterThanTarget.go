package main

import "fmt"

func smallestLetterGreaterThanTarget(letters []rune, target rune) string {
	s := 0
	e := len(letters) - 1
	for s <= e {
		mid := s + (e - s) / 2

		if target < letters[mid] {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return string(letters[s % len(letters)])
}

func main() {
	fmt.Println(smallestLetterGreaterThanTarget([]rune{'c', 'f', 'j', 'k'}, 'j'))
}


