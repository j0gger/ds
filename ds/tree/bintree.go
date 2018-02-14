/*
 * Copyright (c) Sachin Manpathak
 */

// bintree package implements binary tree data structure
package bintree

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

func InorderTraversal(root *Node) []int {
	values := make([]int, 0)
	_Inorder(root, &values)

	return values
}
