package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct {
}

func (c *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, s := range strs {
		builder.WriteString(strconv.Itoa(len(s)))
		builder.WriteByte('#')
		builder.WriteString(s)
	}

	return builder.String()
}

func (c *Solution) Decode(s string) []string {
	var res []string
	n := len(s)
	i := 0

	for i < n {
		j := i
		for j < n && s[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(s[i:j])

		j++

		word := s[j : j+length]
		res = append(res, word)

		i = j + length
	}

	return res
}

func main() {
	solution := Solution{}

	tests := [][]string{
		{},
		{""},
		{"", ""},
		{"lint", "code", "love", "you"},
		{"a#b", "#", "###"},
		{"hola", "  espacios  ", "línea\nnueva"},
	}

	for _, test := range tests {
		encoded := solution.Encode(test)
		decoded := solution.Decode(encoded)
		fmt.Printf("Original: %q\nEncoded: %q\nDecoded: %q\n\n", test, encoded, decoded)
	}
}
