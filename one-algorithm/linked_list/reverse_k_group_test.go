package main

import "github.com/one-go/one-algorithm/internal"

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	head := internal.BuildListNode([]int{1, 2, 3, 4, 5})
	node := reverseKGroup(head, 2)
	node.Print()
}
