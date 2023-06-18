// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150

package main

func maxProfit(stocks []int) int {
	var money, currentPrice, lastPrice int
	hasStock := false

	for i := 0; i < len(stocks); i++ {
		currentPrice = stocks[i]

		if i == 0 {
			continue
		}

		lastPrice = stocks[i-1]

		if currentPrice > lastPrice {
			if !hasStock {
				money -= lastPrice
				hasStock = true
			}
		} else if currentPrice < lastPrice {
			if hasStock {
				money += lastPrice
				hasStock = false
			}
		}
	}
	if hasStock {
		money += stocks[len(stocks)-1]
	}
	return money
}
