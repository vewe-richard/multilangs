package main

import (
	"fmt"
	"io"
	"os"
)

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode

}

type BinaryTree struct{
	root *TreeNode
}

func (bt *BinaryTree)insert(val int){
	if bt.root == nil {
		bt.root = &TreeNode{val,nil,nil}
		return
	}
	root := bt.root
	for root != nil {
		if val < root.val {
			if root.left == nil {
				root.left = &TreeNode{val,nil,nil}
				return
			}
			root = root.left

		} else{
			if root.right == nil {
				root.right = &TreeNode{val,nil,nil}
				return
			}
			root = root.right
		}
	}

}

func print(w io.Writer, node *TreeNode, ns int, ch rune){
	if node == nil {
		return
	}
	fmt.Println(node.val)
	print(w, node.left, ns, ch)
	print(w, node.right, ns, ch)
}

func main() {
	tree := &BinaryTree{}
	tree.insert(100)
	tree.insert(-20)
	tree.insert(-50)
	tree.insert(-15)
	tree.insert(-60)
	tree.insert(50)
	tree.insert(60)
	tree.insert(55)
	tree.insert(85)
	tree.insert(15)
	tree.insert(5)
	tree.insert(-10)
	print(os.Stdout, tree.root, 0, 'M')
}



