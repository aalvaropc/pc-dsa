package main

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range strs {
		count := [26]int{}
		for _, ch := range word {
			count[ch-'a']++
		}

		keyParts := make([]string, 26)
		for i, v := range count {
			keyParts[i] = strconv.Itoa(v)
		}
		key := strings.Join(keyParts, "#")

		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func main() {
	solution := Solution{}
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := solution.GroupAnagrams(words)

	for _, group := range result {
		println(strings.Join(group, ", "))
	}
}
