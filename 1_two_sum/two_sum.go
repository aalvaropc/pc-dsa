package main

import "fmt"

type Solution struct {
}

func (*Solution) calculate(nums []int, target int) []int {
	nmap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, found := nmap[complement]; found {
			return []int{i, j}
		}
		nmap[num] = i
	}
	return []int{}

}

func main() {
	solution := new(Solution)
	nums := []int{2, 7, 11, 15}
	target := 9
	result := solution.calculate(nums, target)
	fmt.Println(result)
}
