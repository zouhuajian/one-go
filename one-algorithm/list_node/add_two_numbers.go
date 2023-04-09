package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeListNode
// create ListNode
func MakeListNode(nums []int) *ListNode {
	dummy := new(ListNode)
	tail := dummy
	for _, num := range nums {
		tail.Next = &ListNode{Val: num}
		tail = tail.Next
	}
	return dummy.Next
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.cn/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	tail := dummy
	last := 0
	for l1 != nil || l2 != nil || last > 0 {
		var v1 = 0
		var v2 = 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + last
		last = sum / 10
		node := &ListNode{Val: sum % 10}
		tail.Next = node
		tail = node
	}
	return dummy.Next
}

func main() {
	arr1 := []int{1, 2, 3}
	list1 := MakeListNode(arr1)
	list1.Print()
	fmt.Println(addTwoNumbers(nil, nil))
}
