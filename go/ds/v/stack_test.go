package v

import (
	"exercise/treenode"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
}

func TestTreeLO(t *testing.T) {
	tree := treenode.GetTreeNode()
	q := Queue[treenode.TreeNode]{}
	q.Offer(*tree)
	for !q.IsEmpty() {
		size := q.Len()
		for i := 0; i < size; i++ {
			root := q.Poll()
			fmt.Print(root.Val)
			if root.Left != nil {
				q.Offer(*root.Left)
			}
			if root.Right != nil {
				q.Offer(*root.Right)
			}
		}
		fmt.Println()
	}
}
