package main

import . "github.com/one-go/one-algorithm/internal"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 排序链表
func sortList2(head *ListNode) *ListNode {
	node := head
	var l = 0
	for node != nil {
		l++
		node = node.Next
	}
	if l == 0 {
		return head
	}
	return sort2(head, 0, l-1)
}

func sort2(node *ListNode, left int, right int) *ListNode {
	if left == right {
		return node
	}
	mid := (left + right) / 2
	subHead1 := node
	for i := left; i < mid; i++ {
		node = node.Next
	}
	subHead2 := node.Next
	// 断开子链表
	node.Next = nil

	subHead1 = sort2(subHead1, left, mid)
	subHead2 = sort2(subHead2, mid+1, right)
	dummy := &ListNode{}
	tail := dummy
	for subHead1 != nil && subHead2 != nil {
		var v1 = subHead1.Val
		var v2 = subHead2.Val
		if v1 < v2 {
			tail.Next = subHead1
			subHead1 = subHead1.Next
		} else {
			tail.Next = subHead2
			subHead2 = subHead2.Next
		}
		tail = tail.Next
	}
	for subHead1 != nil {
		tail.Next = subHead1
		subHead1 = subHead1.Next
		tail = tail.Next
	}
	for subHead2 != nil {
		tail.Next = subHead2
		subHead2 = subHead2.Next
		tail = tail.Next
	}
	return dummy.Next
}
