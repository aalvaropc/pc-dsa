package main

import "fmt"

type Solution struct {
}

func (*Solution) maxProfit(prices []int) int {
	min_prices := int(^uint(0) >> 1)
	max_profit := 0
	for _, price := range prices {
		if price < min_prices {
			min_prices = price
		} else if price-min_prices > max_profit {
			max_profit = price - min_prices
		}
	}

	return max_profit
}

func main() {
	solution := new(Solution)
	prices := []int{7, 1, 5, 3, 6, 4}
	result := solution.maxProfit(prices)
	fmt.Println(result)
}
