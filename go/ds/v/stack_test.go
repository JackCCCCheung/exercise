package v

import (
	"exercise/treenode"
	"fmt"
	"strconv"
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

func TestValid(t *testing.T) {
	str := "{[][]}"
	s := Stack[string]{}
	m := make(map[string]string)
	m[")"] = "("
	m["}"] = "{"
	m["]"] = "["
	isValid := true
	for _, ch := range str {
		item := fmt.Sprintf("%c", ch)
		if e, ok := m[item]; ok {
			if s.IsEmpty() || s.Pop() != e {
				isValid = false
				break
			}
		} else {
			s.Push(item)
		}
	}
	isValid = isValid && s.IsEmpty()
	fmt.Printf("%s is valid : %v\n", str, isValid)
}

// 1+1*2+3/5-2
func TestCalculator(t *testing.T) {
	str := "1+1*2-15/5-2"
	add, minus, multi, div := "+", "-", "*", "/"
	numStack := Stack[int]{}
	tmp := 0
	preOp := add
	for i := 0; i < len(str); i++ {
		item := fmt.Sprintf("%c", str[i])
		isLast, isOp := len(str)-1 == i, str[i] < '0' || str[i] > '9'
		if !isOp {
			if bak, err := strconv.Atoi(item); err == nil {
				tmp = tmp*10 + bak
			} else {
				panic(err)
			}
		}
		if isOp || isLast {
			switch preOp {
			case add:
				numStack.Push(tmp)
				break
			case minus:
				numStack.Push(0 - tmp)
				break
			case multi:
				numStack.Push(numStack.Pop() * tmp)
				break
			case div:
				numStack.Push(numStack.Pop() / tmp)
				break
			}
			preOp = item
			tmp = 0
		}
	}
	result := 0
	for !numStack.IsEmpty() {
		result = result + numStack.Pop()
	}
	fmt.Printf("calculator result: %d \n", result)

}
