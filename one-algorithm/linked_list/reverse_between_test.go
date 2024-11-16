package main

import . "github.com/one-go/one-algorithm/internal"
import (
	"testing"
)

func TestReverseBetween(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})
	head.Print()
	reverseBetween(head, 2, 4)
	head.Print()
}
