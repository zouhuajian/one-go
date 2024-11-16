package main

import . "github.com/one-go/one-algorithm/internal"

/*
环形链表
https://leetcode.cn/problems/linked-list-cycle/
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
