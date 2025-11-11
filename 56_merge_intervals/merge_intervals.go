package main

import (
	"fmt"
	"sort"
)

type Solution struct{}

func (s *Solution) merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
			merged[len(merged)-1] = last
		} else {
			merged = append(merged, interval)
		}
	}
	return merged
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	s := Solution{}
	fmt.Println(s.merge(intervals))
}
