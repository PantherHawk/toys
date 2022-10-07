package main

import "fmt"

func dnaSequence(dna, gene string) int {
	count := 0
	j := 0
	i := 0
	for i < len(dna) {
		if dna[i]  == gene[j] {
			if j == len(gene) - 1 {
				count++
				j = 0
			}
			j++
			i++
		} else {
			j = 0
			i++
		}
	}
	return count
}

func main() {

	fmt.Println(dnaSequence("CATGCATGCATG", "TGC"))
	fmt.Println(dnaSequence("GTTTTAGAAAAG", "TT"))
	fmt.Println(dnaSequence("GTAGAATCATGCA", "GTGTGT"))
}
