package main

import (
	"fmt"
	"math"
	"sort"
)

type Solution struct {
}

func (s *Solution) threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closest := math.MaxInt32

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(target-sum) < abs(target-closest) {
				closest = sum
			}

			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	solution := Solution{}
	fmt.Println(solution.threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
