package main

import "fmt"

type Solution struct{}

func (s Solution) numbersOf1Bits(n uint32) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}

func main() {
	s := Solution{}
	fmt.Println(s.numbersOf1Bits(11))
}
