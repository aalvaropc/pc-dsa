package main

import "fmt"

type Problem struct {
}

func (*Problem) calculate(nums []int, target int) []int {
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
	problem := Problem{}
	nums := []int{2, 7, 11, 15}
	target := 9
	result := problem.calculate(nums, target)
	fmt.Println(result)
}
