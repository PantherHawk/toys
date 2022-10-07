package main

import "fmt"

func dnaSequence(dna, gene string) int {
	count := 0
	geneSequenceLength := len(gene)
	dnaSequenceLength := len(dna)
	lenDiff := dnaSequenceLength - geneSequenceLength

	for i := 0; i <= lenDiff; i++ {
		if dna[i : i + geneSequenceLength] == gene {
			count++
		}
	}
	return count
}

func main() {

	fmt.Println(dnaSequence("CATGCATGCATG", "TGC"))
	fmt.Println(dnaSequence("GTTTTAGAAAAG", "TT"))
	fmt.Println(dnaSequence("GTAGAATCATGCA", "GTGTGT"))
}
