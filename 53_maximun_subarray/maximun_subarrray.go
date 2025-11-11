package main

import "fmt"

type Solution struct{}

func (s *Solution) maximunSubarray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum+nums[i] > nums[i] {
			currentSum = currentSum + nums[i]
		} else {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	s := Solution{}
	fmt.Println(s.maximunSubarray(nums))
}
