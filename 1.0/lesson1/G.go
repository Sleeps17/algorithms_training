package main

import "fmt"

func solutionG() int {
	var n, k, m int
	fmt.Scan(&n, &k, &m)

	if k < m {
		return 0
	}

	cnt := 0
	for n >= k {
		p := n / k
		n %= k
		cnt += p * (k / m)
		n += p * (k % m)
	}
	return cnt
}

func mainG() {
	fmt.Println(solutionG())
}
