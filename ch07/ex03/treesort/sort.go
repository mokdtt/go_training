// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import (
	"bytes"
	"fmt"
	"strings"
)

//!+
type Tree struct {
	value       int
	left, right *Tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = Add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func Add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = Add(t.left, value)
	} else {
		t.right = Add(t.right, value)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return ""
	}
	b := &bytes.Buffer{}
	forEachNode(b, t, preprint)
	return b.String()
}

func forEachNode(b *bytes.Buffer, t *Tree, pre func(b *bytes.Buffer, t *Tree)) {
	if pre != nil {
		pre(b, t)
	}

	if t.left != nil {
		forEachNode(b, t.left, pre)
	}
	if t.right != nil {
		forEachNode(b, t.right, pre)
	}
	depth--
}

var depth int

func preprint(b *bytes.Buffer, t *Tree) {
	blank := 0
	prefix := ""
	if depth != 0 {
		blank = depth*2 - 1
		prefix = "├"
	}
	fmt.Fprintf(b, "%s%s%d\n", prefix, strings.Repeat("─", blank), t.value)
	depth++
}
