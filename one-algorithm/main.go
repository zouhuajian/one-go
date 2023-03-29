package main

import "fmt"

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// create ListNode
func makeListNode(nums [5]int) *ListNode {
	dummy := new(ListNode)
	tail := dummy
	for _, num := range nums {
		tail.Next = &ListNode{Val: num}
		tail = tail.Next
	}
	return dummy.Next
}

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	return nil
}

func main() {
	//var nums = [5]int{1, 2, 3, 4, 5}
	//head := makeListNode(nums)
	//hasCycle := hasCycle(head)
	fmt.Println("Hello world.")

}
