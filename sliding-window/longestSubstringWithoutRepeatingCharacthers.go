package main

// given a string, find the length of the longest substring without repeating characters

import (
	"fmt"
	"math"
)

func lengthOfLongestNonrepeatingSubstring(characters string) int {
	l := 0
	set := make(map[string]bool)
	res := 0.0
	// add characters to accumulation
	// check if dups exist in accumulation
	// while dups exist in accumulation,
	  // slide left up and
	for r, _ := range characters {
		for set[string(characters[r])] {
			delete(set, string(characters[l]))
			l++
		}
		set[string(characters[r])] = true
		res = math.Max(res, float64(r - l + 1))
	}
	return int(res)
}

func main() {
	fmt.Println(lengthOfLongestNonrepeatingSubstring("abcabcbb"))
}
