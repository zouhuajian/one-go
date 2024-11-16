package main

import . "github.com/one-go/one-algorithm/internal"

/*
https://leetcode.cn/problems/reverse-nodes-in-k-group/

给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	fake := &ListNode{}
	fake.Next = head
	pre := fake
	for head != nil {
		for i := 0; i < k-1; i++ {
			head = head.Next
			if head == nil {
				return fake.Next
			}
		}
		subHead := pre.Next
		newSubHead, newSubTail := revers(subHead, head)
		pre.Next = newSubHead
		pre = newSubTail
		head = newSubTail.Next
	}
	return fake.Next
}

/*
*
子组的反转
*/
func revers(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	current := head
	for pre != tail {
		next := current.Next
		current.Next = pre
		pre = current
		// 移动到下一个
		current = next
	}
	return tail, head
}
