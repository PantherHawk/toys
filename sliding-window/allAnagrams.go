package main

import (
	"fmt"
	"reflect"
)

func allAnnagrams(s, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	sCount, pCount := make(map[string]int), make(map[string]int)
	l := 0
	for i, v:= range p {
		pCount[string(v)] = pCount[string(v)] + 1
		sCount[string(s[i])] = sCount[string(s[i])] + 1
	}
	fmt.Println(sCount)
	res := []int{0}
	for r := len(p); r < len(s); r++ {
		fmt.Println("r: ", r)
		sCount[string(s[r])] = sCount[string(s[r])] + 1
		sCount[string(s[l])] = sCount[string(s[l])] - 1

		if sCount[string(s[l])] == 0 {
			fmt.Print("here...\n")
			delete(sCount, string(s[l]))
		}
		l++
		fmt.Printf("sCount: %v\n", sCount)
		if reflect.DeepEqual(sCount, pCount) {
			res = append(res, l)
		}
	}
	return res
}


func main() {
	fmt.Println(allAnnagrams("cbaebabacd", "abc"))
}
