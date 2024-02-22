package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainB() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	arr := make([]int64, n)
	for i := range arr {
		fmt.Fscan(rd, &arr[i])
	}

	pref := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		pref[i] = pref[i-1] + arr[i-1]
	}

	_min := int64(0)
	ans := pref[1]

	for i := 1; i <= n; i++ {
		ans = max(ans, pref[i]-_min)

		_min = min(_min, pref[i])
	}

	fmt.Println(ans)
}
