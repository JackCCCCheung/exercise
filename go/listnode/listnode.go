package listnode

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) toString() string {
	result := ""

	tmp := l
	for tmp != nil {
		result = result + "->" + strconv.Itoa(tmp.Val)
		tmp = tmp.Next
	}
	return result
}

func getListNode() *ListNode {
	result := ListNode{Val: 1, Next: nil}
	result.Next = &ListNode{2, nil}
	result.Next.Next = &ListNode{3, nil}
	result.Next.Next.Next = &ListNode{4, nil}
	result.Next.Next.Next.Next = &ListNode{5, nil}
	return &result
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{Next: head, Val: -1}
	slow, fast := result, head
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return result.Next
}
