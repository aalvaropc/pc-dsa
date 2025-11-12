package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

func (s *Solution) longestIncreasingSubsequence(nums []int) int {
	sub := []int{}

	for _, x := range nums {
		i := sort.SearchInts(sub, x)
		if i == len(sub) {
			sub = append(sub, x)
		} else {
			sub[i] = x
		}
	}

	return len(sub)
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	s := Solution{}
	fmt.Println(s.longestIncreasingSubsequence(nums))
}
