package main

import (
	"bufio"
	"fmt"
	"os"
)

type GeneralTree struct {
	root string
	m    map[string][]string
}

func (tree *GeneralTree) search(root string, x string) bool {
	if root == "" {
		return false
	}

	if root == x {
		return true
	}

	for _, son := range tree.m[root] {
		if tree.search(son, x) {
			return true
		}
	}

	return false
}

func (tree *GeneralTree) Add(parent, son string) {
	if tree.m == nil {
		tree.m = make(map[string][]string)
	}
	tree.m[parent] = append(tree.m[parent], son)
}

func (tree *GeneralTree) Search(x string) bool {
	return tree.search(tree.root, x)
}

func mainB() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	var tree GeneralTree

	for i := 0; i < n-1; i++ {
		var son, parent string
		fmt.Fscan(rd, &son, &parent)

		if i == 0 {
			tree.root = parent
		}

		tree.Add(parent, son)
	}

	for {

		var first, second string

		if n, _ := fmt.Fscan(rd, &first, &second); n != 2 {
			break
		}

		if tree.search(first, second) {
			fmt.Print(1, " ")
		} else if tree.search(second, first) {
			fmt.Print(2, " ")
		} else {
			fmt.Print(0, " ")
		}
	}
	fmt.Println()
}
