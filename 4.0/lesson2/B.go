package main

import "fmt"

func isEqual(from1, from2, long int, h []int, x []int, P int) bool {
	h1 := h[from1+long-1] - (h[from1-1]*x[long])%P
	h2 := h[from2+long-1] - (h[from2-1]*x[long])%P

	if h1 < 0 {
		h1 += P
	}

	if h2 < 0 {
		h2 += P
	}

	//fmt.Println("h1: ", h1)
	//fmt.Println("h2: ", h2)

	return h1 == h2
}

func mainB() {
	var s string
	fmt.Scan(&s)
	X := 257
	P := 1000000007

	n := len(s)
	h := make([]int, n+1)
	x := make([]int, n+1)
	x[0] = 1
	for i := 1; i <= n; i++ {
		h[i] = (h[i-1]*X + int(s[i-1])) % P
		x[i] = (x[i-1] * X) % P
	}
	maxEqualLen := 0
	for curEqualLen := 1; curEqualLen < n; curEqualLen++ {
		if isEqual(1, n+1-curEqualLen, curEqualLen, h, x, P) {
			if curEqualLen > maxEqualLen {
				maxEqualLen = curEqualLen
			}
		}
	}
	if maxEqualLen == 0 {
		fmt.Println(n)
	} else {
		fmt.Println(n - maxEqualLen)
	}
}
