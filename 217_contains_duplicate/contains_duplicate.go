package main

import "fmt"

type Solution struct {
}

func (*Solution) containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, n := range nums {
		if seen[n] {
			return true
		}
		seen[n] = true
	}

	return false
}

func main() {
	solution := new(Solution)
	nums := []int{1, 2, 3, 1}
	result := solution.containsDuplicate(nums)
	fmt.Println(result)

}
