package internal

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Print() {
	h := head
	for h != nil {
		fmt.Printf("%v", h.Val)
		if h.Next != nil {
			fmt.Print("->")
		}
		h = h.Next
	}
	fmt.Println()
}

func BuildListNode(nums []int) *ListNode {
	fakeHead := new(ListNode)
	tail := fakeHead
	for _, num := range nums {
		tail.Next = &ListNode{Val: num}
		tail = tail.Next
	}
	return fakeHead.Next
}
