package main

import (
	"fmt"
)

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	var l, r int
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func reverse(arr []int) []int {
	n := len(arr)
	rev := make([]int, n)

	for i, elem := range arr {
		rev[n-1-i] = elem
	}

	return rev
}

func isEqualD(from1, from2, long int, h, x []int, P int) bool {
	if long == 0 {
		return true
	}
	h1 := h[from1+long] - (h[from1]*x[long])%P
	h2 := h[from2+long] - (h[from2]*x[long])%P

	if h1 < 0 {
		h1 += P
	}

	if h2 < 0 {
		h2 += P
	}

	fmt.Println("h1: ", h1)
	fmt.Println("h2: ", h2)

	return h1 == h2
}

func mainD() {

	X := 257
	P := 1000000007

	var n, m int
	fmt.Scan(&n, &m)

	cubes := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cubes[i])
	}

	rev := reverse(cubes)
	cubes = append(cubes, rev...)

	h := make([]int, 2*n+1)
	x := make([]int, 2*n+1)
	x[0] = 1

	for i := 1; i < 2*n; i++ {
		h[i] = (h[i-1]*X + cubes[i]) % P
	}

	for i := 1; i < 2*n+1; i++ {
		x[i] = (x[i-1] * X) % P
	}

	ans := make([]int, 0)
	for i := n / 2; i >= 0; i-- {
		if isEqualD(2*n-i, i+1, i, h, x, P) {
			ans = append(ans, n-i)
		}
	}

	for i := 0; i < len(ans); i++ {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Println()
}
