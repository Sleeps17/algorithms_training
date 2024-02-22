package main

import (
	"fmt"
	"math"
)

func mainA() {
	var n, m int
	fmt.Scan(&n, &m)

	seq := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&seq[i])
	}

	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scan(&l, &r)

		minim := math.MaxInt
		maxim := math.MinInt

		for j := l; j <= r; j++ {
			minim = min(minim, seq[j])
			maxim = max(maxim, seq[j])
		}

		if minim == maxim {
			fmt.Println("NOT FOUND")
		} else {
			fmt.Println(maxim)
		}
	}
}
