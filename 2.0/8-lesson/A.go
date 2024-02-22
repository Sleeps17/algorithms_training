package main

import (
	"bufio"
	"fmt"
	"os"
)

type BNode struct {
	Value int
	Left  *BNode
	Right *BNode
}

type BinaryTree struct {
	Root *BNode
}

func NewBNode(Value int) *BNode {
	return &BNode{
		Value: Value,
	}
}

func (tree *BinaryTree) add(root **BNode, x int) bool {
	if *root == nil {
		*root = NewBNode(x)
		return true
	}

	if x < (*root).Value {
		return tree.add(&(*root).Left, x)
	} else if x > (*root).Value {
		return tree.add(&(*root).Right, x)
	}

	return false
}

func (tree *BinaryTree) search(root *BNode, x int) bool {
	if root == nil {
		return false
	}

	if x < root.Value {
		return tree.search(root.Left, x)
	} else if x > root.Value {
		return tree.search(root.Right, x)
	}

	return true
}

func (tree *BinaryTree) print(root *BNode, depth int) {
	if root == nil {
		return
	}

	tree.print(root.Left, depth+1)

	for i := 0; i < depth; i++ {
		fmt.Print(".")
	}
	fmt.Println(root.Value)

	tree.print(root.Right, depth+1)
}

func (tree *BinaryTree) Add(x int) bool {
	return tree.add(&tree.Root, x)
}

func (tree *BinaryTree) Search(x int) bool {
	return tree.search(tree.Root, x)
}

func (tree *BinaryTree) Print() {
	tree.print(tree.Root, 0)
}

func mainA() {
	rd := bufio.NewReader(os.Stdin)

	var tree BinaryTree

	for {
		var action string

		if n, _ := fmt.Fscan(rd, &action); n != 1 {
			break
		}

		switch action {
		case "ADD":
			var x int
			fmt.Fscan(rd, &x)
			f := tree.Add(x)
			if f {
				fmt.Println("DONE")
			} else {
				fmt.Println("ALREADY")
			}
		case "SEARCH":
			var x int
			fmt.Fscan(rd, &x)
			f := tree.Search(x)
			if f {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case "PRINTTREE":
			tree.Print()
		}
	}
}
