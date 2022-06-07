package DataStructures

import (
	"testing"
)

func TestInsertLeft(t *testing.T) {
	n := &BinaryTree{Key: 25}

	n.Insert(10)

	if n.Left != nil && n.Left.Key == 10 {
		return
	}

	t.Fatalf(`Binary Tree Insert = %d want 10`, n.Left.Key)
}

func TestInsertRight(t *testing.T) {
	n := &BinaryTree{Key: 25}

	n.Insert(30)

	if n.Right != nil && n.Right.Key == 30 {
		return
	}

	t.Fatalf(`Binary Tree Insert = %v want 10`, n.Right)
}

func TestSearch(t *testing.T) {
	n := &BinaryTree{Key: 25}

	keys := [10]int{19, 60, 39, 23, 29, 67, 20, 5, 4, 30}

	for _, k := range keys {
		n.Insert(k)
	}

	found := n.Search(29)

	if !found {
		t.Fatalf("Binary Tree Search = %v want true", found)
	}
}
