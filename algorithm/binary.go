package algorithm

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func render(w io.Writer, node *BinaryNode, ns int, ch string) {
	if node != nil {
		for i := 0; i < ns; i++ {
			fmt.Fprint(w, " ")
		}
		fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	}
	render(w, node.left, ns+2, "L")
	render(w, node.right, ns+2, "R")
}
