package main


import (
	"fmt"
	"strings"
)

func mirrorReverse(a []string, s int, e int) {
	for s < e {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}
}

func reverseWords(a []string) []string {
	mirrorReverse(a, 0, len(a) - 1)
	
	wordStart := -1
	for i, v := range(a) {
		if len(strings.TrimSpace(v)) == 0 {
			if wordStart > -1 {
				mirrorReverse(a, wordStart, i - 1)
				wordStart = -1
			}
		} else if i == len(a) - 1 {
			if wordStart > -1 {
				mirrorReverse(a, wordStart, i)
			}
		} else {
			if wordStart < 0 {
				wordStart = i
			}
		}
	}
	return a
}

func main() {
	fmt.Println(reverseWords([]string{
                "p", "e", "r", "f", "e", "c", "t", " ", " ",
                "m", "a", "k", "e", "s", " ",
                "p", "r", "a", "c", "t", "i", "c", "e" }))
	}
