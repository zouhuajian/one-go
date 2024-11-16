package main

import "github.com/one-go/one-algorithm/internal"

import (
	"testing"
)

func TestSortList(t *testing.T) {
	head := internal.BuildListNode([]int{4, 2, 1, 3})
	node := sortList(head)
	node.Print()
}
