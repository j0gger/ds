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
	sortedVals := InorderTraversal(root)
	assert.Equal(t, vals, sortedVals)
	
}
