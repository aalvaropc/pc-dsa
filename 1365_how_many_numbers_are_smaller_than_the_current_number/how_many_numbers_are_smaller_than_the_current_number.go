package main

import "fmt"

type Solution struct{}

func (s *Solution) smallerNumbersThanCurrent(nums []int) []int {
	count := make([]int, 101)

	for _, num := range nums {
		count[num]++
	}

	for i := 1; i < 101; i++ {
		count[i] += count[i-1]
	}

	result := make([]int, len(nums))
	for i, num := range nums {
		if num == 0 {
			result[i] = 0
		} else {
			result[i] = count[num-1]
		}
	}
	return result
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	s := Solution{}
	fmt.Println(s.smallerNumbersThanCurrent(nums))
}
