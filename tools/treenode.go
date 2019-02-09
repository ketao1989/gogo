package tools

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func BuildTree() *TreeNode {

	left := &TreeNode{val: 2, left: nil, right: nil}
	right := &TreeNode{val: 3, left: nil, right: nil}

	head := &TreeNode{val: 1, left: left, right: right}

	return head
}

func (node *TreeNode) GetLeft() *TreeNode {

	return node.left
}

func (node *TreeNode) GetRight() *TreeNode {

	return node.right
}

func (node *TreeNode) GetVal() int {

	return node.val
}

func (node *TreeNode) PrintTree() {

	var tmp *TreeNode = node
	if tmp != nil {
		fmt.Print(" ", tmp.val)
	}

	if tmp.left != nil {
		tmp.left.PrintTree()
	}

	if tmp.right != nil {
		tmp.right.PrintTree()
	}
}
