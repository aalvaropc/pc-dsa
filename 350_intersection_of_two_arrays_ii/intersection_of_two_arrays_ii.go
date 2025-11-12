package main

import "fmt"

type Solution struct{}

func (s *Solution) intersectionTwoArraysII(nums1 []int, nums2 []int) []int {
	count := make(map[int]int)
	for _, num := range nums1 {
		count[num]++
	}

	var result []int
	for _, num := range nums2 {
		if count[num] > 0 {
			result = append(result, num)
			count[num]--
		}
	}
	return result
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	s := Solution{}
	fmt.Println(s.intersectionTwoArraysII(nums1, nums2))
}
