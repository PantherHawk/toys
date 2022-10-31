package main

import (
	"fmt"
	"math"
)

// given two strings s and t return the minimum window in s 
// which will contain all the characters in t. 
// If no such window in s exists that 
// covers all characters in t, return empty string.

func minWindowSubstring(s, t string) string {
	// handle empty target string edge case
	if len(t) < 1 {
		return ""
	}
	// initialize two maps, 1) for keepin track of found target letters, and 2) for keeping track of the sliding window
	countT, window := make(map[string]int), make(map[string]int)
	// fill up target letters map with targets from t
	for _, v := range t {
		countT[string(v)] = 1 + countT[string(v)]
	}
	// keep track of required letter count with two pointers
	have, need := 0, len(countT)
	// keep track of indeces of sliding window for result, and initialize length to very big
	res, resLen := make([]int, 2), math.Inf(1)
	l := 0
	// iterate over s
	for i, _ := range s {
		c := string(s[i])
		// expand window to the right
		window[c] = 1 + window[c]
		// check if current letter is a target letter
		if countT[c] > 0 && window[c] == countT[c] {
			have++
		}
		// if we have all the letters we need
		for have == need {
			// start minimizing window from the left
			// check if we have a smaller window and save the indeces and new length for response
			if float64(i - l + 1) < resLen {
				res[0], res[1] = l, i
				resLen = float64(i - l + 1)
			}
			// pop from left of window
			window[string(s[l])] = window[string(s[l])] - 1
			// check if we have removed a target letter
			d := string(s[l])
			if countT[d] > 0 && window[d] < countT[d] {
				have--
			}
			l++
		}
	}
	l, r := res[0], res[1]
	if resLen > 0.0 {
		return string(s[l:r+1])
	} 
	return ""
}

func main() {
	fmt.Println(minWindowSubstring("ADOBECODEBANC", "ABC"))
}
