package node

import "strconv"

type ListNode struct {
	Next *ListNode
	Val  int
}

func (l *ListNode) String() string {
	result := ""
	tmp := l
	for tmp != nil {
		result = result + "->" + strconv.Itoa(tmp.Val)
		tmp = tmp.Next
	}
	return result
}

type DListNode struct {
	Next *DListNode
	Prev *DListNode
	Val  int
}

//func (l *DListNode) String() string {
//	result := ""
//	tmp := l
//	for tmp != nil {
//	}
//	return result
//}
