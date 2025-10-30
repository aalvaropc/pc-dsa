package main

import "fmt"

type Solution struct {
}

func (solution *Solution) is_anagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[rune]int)

	for _, ch := range s {
		counter[ch]++
	}

	for _, ch := range t {
		if counter[ch] == 0 {
			return false
		}
		counter[ch]--
	}

	return true

}

func main() {
	solution := Solution{}
	s := "anagram"
	t := "nagaram"

	result := solution.is_anagram(s, t)
	fmt.Println(result)
}
