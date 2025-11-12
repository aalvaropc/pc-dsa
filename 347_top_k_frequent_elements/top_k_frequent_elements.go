package main

import "fmt"

type Solution struct{}

func (s Solution) topKFrequentElements(nums []int, k int) []int {

	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	freq := make([][]int, len(nums)+1)
	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	res := []int{}
	for i := len(freq) - 1; i >= 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}

func main() {
	s := Solution{}
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(s.topKFrequentElements(nums, k))
}
