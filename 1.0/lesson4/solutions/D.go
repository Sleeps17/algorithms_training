package main

import "fmt"

func mainD() {
	var n int
	fmt.Scan(&n)

	dict := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var cnt int
		fmt.Scan(&cnt)
		dict[i] = cnt
	}

	var k int
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		var p int
		fmt.Scan(&p)
		dict[p]--
	}

	for i := 1; i <= n; i++ {
		if dict[i] < 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
