package main

import "fmt"

type Solution struct{}

func (s *Solution) longestConsecutiveSequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for num := range numSet {
		if !numSet[num-1] {
			length := 1
			for numSet[num+length] {
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	s := Solution{}
	fmt.Println(s.longestConsecutiveSequence(nums))
}
