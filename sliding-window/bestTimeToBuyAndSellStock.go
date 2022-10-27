package main

import "fmt"
// given an array, each value represents the stock's value at a particular day
// we can buy on one day and sell on another day
// return the profit

func buyLowSellHight(prices []int) int {
	l, r := 0, 1
	maxProfit := 0

	for r < len(prices) {
		// profitable?
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			l = r
		}
		r++
	}
	return maxProfit
}

func main() {
	fmt.Println(buyLowSellHight([]int{7, 1, 5, 3, 6, 4}))
}
