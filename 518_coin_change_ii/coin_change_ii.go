package main

import "fmt"

type Solution struct{}

func (s *Solution) change_change_ii(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for x := coin; x <= amount; x++ {
			dp[x] += dp[x-coin]
		}
	}
	return dp[amount]
}

func main() {
	s := Solution{}
	fmt.Println(s.change_change_ii(5, []int{1, 2, 5}))
}
