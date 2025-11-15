package main

import "fmt"

type Solution struct {
}

func (s *Solution) reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		res = (res << 1) | (num & 1)
		num >>= 1
	}
	return res
}

func main() {
	s := Solution{}
	var n uint32 = 43261596
	fmt.Println(s.reverseBits(n))
}
