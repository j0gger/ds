/*
 * Copyright (c) Sachin Manpathak
 */

// bintree package implements binary tree data structure
package bintree

import (
	"errors"
)

type Node struct {
	val int
	left *Node
	right *Node
}


func Insert(root *Node, val int) *Node {
	if root == nil {
		newNode := Node{val, nil, nil}
		return &newNode
	}

	if val < root.val {
		root.left = Insert(root.left, val)
		return root
	}

	root.right = Insert(root.right, val)

	return root
}

func _Inorder(root *Node, values *[]int) {
	if root == nil {
		return
	}

	_Inorder(root.left, values)
	*values = append(*values, root.val)
	_Inorder(root.right, values)
}

func TreeToInorder(root *Node) []int {
	values := make([]int, 0)
	_Inorder(root, &values)

	return values
}

func GetSmallerNeighbor(root *Node, val int) (*Node,  error) {
	if root == nil {
		return root, errors.New("Not found")
	}
	
	if root.val <= val {
		return root, nil
	}

	return GetSmallerNeighbor(root.left, val)
}

func _Abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func _Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func _IsBalanced(root *Node) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftOk, leftHeight := _IsBalanced(root.left)

	if leftOk == false {
		return false, 0
	}

	rightOk, rightHeight := _IsBalanced(root.right)

	if rightOk == false {
		return false, 0
	}

	return _Abs(leftHeight - rightHeight) <= 1, _Max(leftHeight, rightHeight) + 1
}

func IsBalanced(root *Node) bool {
	isBalanced, _ := _IsBalanced(root)
	return isBalanced
}
