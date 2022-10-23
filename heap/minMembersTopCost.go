package main

import (
	"fmt"
	"sort"
)

type visit struct {
	Cost int
}

type member struct {
	Id string
	Visits []visit
}

func (m member) getTotalCost() int {
	cost := 0
	for _, v := range(m.Visits) {
		cost += v.Cost
	}
	return cost
}

func getMembersForThreshold(members []member, threshold float64) []member {
	
	result := []member{}

	// create map of member index in list -> member totalcost
	memberCosts := make(map[int]int)
	totalCost := 0
	for i, m := range members {
		memberCosts[i] = m.getTotalCost()
		totalCost += memberCosts[i]
	}
	fmt.Println(memberCosts)
	// sort map by member totalcost
	keys := make([]int, 0, len(memberCosts))
	
	for key := range memberCosts {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return memberCosts[keys[i]] > memberCosts[keys[j]]
	})

	// get targetcost form totalcost x threshold
	targetCost := int(threshold * float64(totalCost))
	// iterate over map until targetcost < 0
	for _, idx := range keys {
		if targetCost > 0 {
			result = append(result, members[idx])
			targetCost -= memberCosts[idx]
		}
	}
	return result
}

func main() {
	a := visit{5}
	b := visit{10}
	c := visit{5}

	firstVisits := []visit{a, b}
	secondVisits := []visit{c}

	first := member{"first", firstVisits}
	second := member{"second", secondVisits}

	members := []member{first, second}

	//given a list of members and a threshold cost, determine which members make up the threshold cost with their visits

	threshold := 0.75

	fmt.Println(getMembersForThreshold(members, threshold))
}
