package main

import . "github.com/one-go/one-algorithm/internal"

/*
23. 合并 K 个升序链表
https://leetcode.cn/problems/merge-k-sorted-lists/

思路：分治合并
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := (left + right) >> 1
	list1 := merge(lists, left, mid)
	list2 := merge(lists, mid+1, right)
	return mergeTwo(list1, list2)
}

func mergeTwo(list1 *ListNode, list2 *ListNode) *ListNode {
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
