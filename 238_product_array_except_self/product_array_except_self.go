package main

import "fmt"

type Solution struct {
}

func (*Solution) productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return answer
}

func main() {
	solution := new(Solution)
	nums := []int{1, 2, 3, 4}
	result := solution.productExceptSelf(nums)
	fmt.Println(result)
}
