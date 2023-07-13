package main

import "fmt"

func maxJ(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minJ(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mainH() {
	var a, b, n, m int
	fmt.Scan(&a, &b, &n, &m)

	min1 := n + (n-1)*a
	max1 := n + (n+1)*a
	min2 := m + (m-1)*b
	max2 := m + (m+1)*b

	if min2 > max1 || min1 > max2 {
		fmt.Println(-1)
	} else {
		fmt.Println(maxJ(min1, min2), minJ(max1, max2))
	}
}
