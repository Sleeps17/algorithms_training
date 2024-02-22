package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func mainA() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(rd, &arr[i])
	}

	slices.Sort(arr)

	var k int
	fmt.Fscan(rd, &k)

	for i := 0; i < k; i++ {
		var l, r int
		fmt.Fscan(rd, &l, &r)

		fmt.Print(upper_bound(arr, r)-lower_bound(arr, l), " ")
	}
	fmt.Println()
}
