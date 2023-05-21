package treenode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetTreeNode() *TreeNode {
	result := TreeNode{Val: 5, Left: nil, Right: nil}
	result.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	result.Right = &TreeNode{Val: 6, Left: nil, Right: nil}
	result.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	result.Left.Right = &TreeNode{Val: 1, Left: nil, Right: nil}
	result.Right.Left = &TreeNode{Val: 7, Left: nil, Right: nil}
	result.Right.Right = &TreeNode{Val: 8, Left: nil, Right: nil}
	return &result
}
