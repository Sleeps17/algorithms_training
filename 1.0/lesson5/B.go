package main

import "fmt"

func mainB() {
	var n int
	var k int64
	fmt.Scan(&n, &k)

	dict := make(map[int64]int)

	pref := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		var num int64
		fmt.Scan(&num)
		pref[i] = pref[i-1] + num
	}

	ans := 0

	for i := 0; i <= n; i++ {
		dict[pref[i]]++

		ans += dict[pref[i]-k]
	}

	fmt.Println(ans)
}
