package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct{}

func (s *Solution) reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	s := Solution{}
	reversed := s.reverseLinkedList(head)

	var result []int
	for reversed != nil {
		result = append(result, reversed.Val)
		reversed = reversed.Next
	}
	fmt.Println(result)
}
