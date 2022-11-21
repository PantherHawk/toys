package main

import "fmt"

func recursePattern(r, c int) {
	if r == 0 {
		return
	}
	for c < r {
		fmt.Print("*")
		c = c + 1
	}
	fmt.Println()
	recursePattern(r - 1, 0)
}

func main() {
	recursePattern(4, 0)
}
