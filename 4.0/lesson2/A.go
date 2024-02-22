package main

import (
	"fmt"
)

func mainA() {
	var s string
	fmt.Scan(&s)

	var q int
	fmt.Scan(&q)

	requests := make([][3]int, q)
	for i := 0; i < q; i++ {
		var l, a, b int
		fmt.Scan(&l, &a, &b)
		requests[i] = [3]int{l, a, b}
	}

	X := 257
	P := 1000000007

	isEqual := func(from1, from2, long int, h, x []int) bool {
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

	n := len(s)
	h := make([]int, n+1)
	x := make([]int, n+1)

	x[0] = 1
	h[0] = 0

	for i := 1; i <= n; i++ {
		h[i] = (h[i-1]*X + int(s[i-1]-'a') + 1) % P
		x[i] = (x[i-1] * X) % P
	}

	for _, req := range requests {
		l, a, b := req[0], req[1], req[2]
		if isEqual(a+1, b+1, l, h, x) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
