package main

import "fmt"

type Solution struct{}

func (s *Solution) findMinimumRotatedSortedArray(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	s := Solution{}
	fmt.Println(s.findMinimumRotatedSortedArray(nums))
}
