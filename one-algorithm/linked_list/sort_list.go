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
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow := head
	fast := head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	subHead1 := sort(head, mid)
	subHead2 := sort(mid, tail)
	return mergeList(subHead1, subHead2)
}

func mergeList(head1 *ListNode, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for head1 != nil && head2 != nil {
		var v1 = head1.Val
		var v2 = head2.Val
		if v1 < v2 {
			tail.Next = head1
			head1 = head1.Next
		} else {
			tail.Next = head2
			head2 = head2.Next
		}
		tail = tail.Next
	}
	for head1 != nil {
		tail.Next = head1
		head1 = head1.Next
		tail = tail.Next
	}
	for head2 != nil {
		tail.Next = head2
		head2 = head2.Next
		tail = tail.Next
	}
	return dummy.Next
}
