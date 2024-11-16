package main

import . "github.com/one-go/one-algorithm/internal"

/*
21. 合并两个有序链表
https://leetcode.cn/problems/merge-two-sorted-lists/
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	fake := &ListNode{}
	tail := fake
	for list1 != nil && list2 != nil {
		v1 := list1.Val
		v2 := list2.Val
		if v1 >= v2 {
			tail.Next = list2
			list2 = list2.Next
		} else {
			tail.Next = list1
			list1 = list1.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	}

	if list2 != nil {
		tail.Next = list2
	}

	return fake.Next
}
