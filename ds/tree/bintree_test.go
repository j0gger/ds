/*
 * Copyright(c) 2018, Sachin Manpathak
 */

package bintree

import (
	"sort"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	vals := []int{45, 12, 33, 1, 4, 15, 11}

	var root *Node

	for _, v := range vals {
		root = Insert(root, v)
	}

	sort.Ints(vals)
	sortedVals := TreeToInorder(root)
	assert.Equal(t, vals, sortedVals)
	
}

func TestUnBalanced(t *testing.T) {
	vals := []int{67, 32, 11, 0, 3, 12, 16, 42, 39}
	var root *Node
	for _, v := range vals {
		root = Insert(root, v)
	}

	assert.Equal(t, false, IsBalanced(root))
}

func TestUnBalanced2(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5, 6}
	var root *Node
	for _, v := range vals {
		root = Insert(root, v)
	}

	assert.Equal(t, false, IsBalanced(root))
}


func TestBalanced(t *testing.T) {
	vals := []int{50, 25, 75, 30, 20, 70, 80, 35, 28}
	var root *Node
	for _, v := range vals {
		root = Insert(root, v)
	}

	assert.Equal(t, true, IsBalanced(root))	
}
