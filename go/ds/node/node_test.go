package node

import (
	"fmt"
	"testing"
)

func TestListNode(t *testing.T) {
	l := buildNode()
	fmt.Println(l.String())
}

func TestDelNode(t *testing.T) {
	l := buildNode()
	delIndex := 5
	fast, slow := l, l
	for fast != nil && delIndex > 0 {
		delIndex--
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	fmt.Println(l.String())
}

func TestDNode(t *testing.T) {
	fmt.Println(buildDNode())
}

func buildNode() *ListNode {
	return &ListNode{
		Val: 0, Next: &ListNode{
			Val: 1, Next: &ListNode{
				Val: 2, Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
	}
}

func buildDNode() *DListNode {
	head := &DListNode{Val: 0}
	second := &DListNode{Val: 1, Prev: head}
	head.Next = second
	thirdly := &DListNode{Val: 2, Prev: second}
	second.Next = thirdly
	fourthly := &DListNode{Val: 3, Prev: thirdly}
	thirdly.Next = fourthly
	fifth := &DListNode{Val: 4, Prev: fourthly}
	fourthly.Next = fifth
	return head
}
