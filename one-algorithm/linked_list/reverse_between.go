package main

import . "github.com/one-go/one-algorithm/internal"

/*
*
  - 反转链表2
  - https://leetcode.cn/problems/reverse-linked-list-ii/description/
  - @param

1,2,3,4,5
2->4, 3->2, 1->3
1,3,2,4,5
3->5, 4->3, 1->4
  - @return 1,4,3,2,5
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	fake := &ListNode{Val: -1}
	fake.Next = head
	pre := fake
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	current := pre.Next
	for i := 0; i < right-left; i++ {
		next := current.Next
		// 连接尾部未反转的部分
		current.Next = next.Next
		// 反转，注意反转时使用 pre.Next 而不是current，因为pre.Next会随着反转而变动
		next.Next = pre.Next
		// 连接头部
		pre.Next = next
	}
	return fake.Next
}
