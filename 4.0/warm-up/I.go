package main

import "fmt"

func mainI() {
	var seq string
	fmt.Scan(&seq)
	result := isBalanced(seq)
	if result {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func isBalanced(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (char == ')' && top == '(') || (char == ']' && top == '[') || (char == '}' && top == '{') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
