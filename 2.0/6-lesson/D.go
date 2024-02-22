package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func mainD() {
	rd := bufio.NewReader(os.Stdin)

	var a, k, b, m, x int64
	fmt.Fscan(rd, &a, &k, &b, &m, &x)

	l := int64(0)
	r := int64(math.MaxInt64 / max(a, b))

	for l < r {
		mid := (l + r) / 2

		var n1 int64 = a*mid - a*(mid/k)
		var n2 int64 = b*mid - b*(mid/m)

		var res int64 = n1 + n2

		if res < x {
			l = mid + 1
		} else {
			r = mid
		}
	}

	fmt.Println(l)
}
