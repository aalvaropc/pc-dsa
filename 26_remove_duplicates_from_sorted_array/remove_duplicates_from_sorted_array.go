package main

import "fmt"

type Solution struct{}

func (s *Solution) removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	s := Solution{}
	k := s.removeDuplicates(nums)
	fmt.Println(k, nums[:k])
}
