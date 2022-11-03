package main

import "fmt"

func ntoBoth(n int) {
	if n == 1 {
		fmt.Println(n)
		return
	}
	fmt.Println(n)
	ntoBoth(n - 1)
	fmt.Println(n)
}

func main() {
	ntoBoth(5)
}
