package main

import (
	"bufio"
	"fmt"
	"os"
)

func countPalindromicSubstrings(s string) int {
	n := len(s)
	count := 0

	for i := 0; i < 2*n-1; i++ {
		left := i / 2
		right := left + i%2
		for left >= 0 && right < n && s[left] == s[right] {
			count++
			left--
			right++
		}
	}

	return count
}

func mainE() {
	var input string
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscan(rd, &input)
	result := countPalindromicSubstrings(input)
	fmt.Println(result)
}
