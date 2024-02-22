package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainD() {
	rd := bufio.NewReader(os.Stdin)

	var l, k int
	fmt.Fscan(rd, &l, &k)

	arr := make([]int, k)
	for i := range arr {
		fmt.Fscan(rd, &arr[i])
	}

	s := 0
	f := k - 1

	mid := (l - 1) / 2

	for i := 0; i < k; i++ {
		if arr[i] > mid {
			break
		}

		s = i
	}

	mid = l / 2
	for i := k - 1; i >= 0; i-- {
		if arr[i] < mid {
			break
		}

		f = i
	}

	for i := s; i <= f; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
