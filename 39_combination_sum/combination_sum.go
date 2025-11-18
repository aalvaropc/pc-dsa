package main

import "fmt"

type Solution struct{}

func (s *Solution) combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	path := []int{}

	var backtrack func(start int, remaining int)
	backtrack = func(start int, remaining int) {
		if remaining == 0 {
			comb := make([]int, len(path))
			copy(comb, path)
			res = append(res, comb)
			return
		}
		if remaining < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, remaining-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return res
}

func main() {
	solution := &Solution{}
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := solution.combinationSum(candidates, target)
	fmt.Println(result)
}
