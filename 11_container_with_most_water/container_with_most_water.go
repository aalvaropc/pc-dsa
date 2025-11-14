package main

import "fmt"

type Solution struct{}

func (s *Solution) maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left
		h := min(height[left], height[right])
		area := width * h
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	s := Solution{}
	fmt.Println(s.maxArea(nums))
}
