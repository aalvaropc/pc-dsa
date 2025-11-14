package main

import "fmt"

type Solution struct{}

func (s *Solution) maximumProductSubarray(nums []int) int {
	maxProd, minProd, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			maxProd, minProd = minProd, maxProd
		}

		maxProd = max(num, num*maxProd)
		minProd = min(num, num*minProd)
		result = max(result, maxProd)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, -2, 4}
	s := Solution{}
	fmt.Println(s.maximumProductSubarray(nums))
}
