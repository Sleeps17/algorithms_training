package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainA() {
	rd := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(rd, &n, &q)

	arr := make([]int64, n)
	for i := range arr {
		fmt.Fscan(rd, &arr[i])
	}

	pref := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		pref[i] = pref[i-1] + arr[i-1]
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(rd, &l, &r)

		fmt.Println(pref[r] - pref[l-1])
	}
}
