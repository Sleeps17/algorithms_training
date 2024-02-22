package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var a, b, c, d float64

const eps float64 = 1e-5

func f(x float64) float64 {
	return a*math.Pow(x, 3) + b*math.Pow(x, 2) + c*x + d
}

func mainC() {
	rd := bufio.NewReader(os.Stdin)

	fmt.Fscan(rd, &a, &b, &c, &d)

	if a < 0 {
		a *= -1
		b *= -1
		c *= -1
		d *= -1
	}

	var l, r float64 = -2000.0, 2000.0

	for r-l >= eps {
		m := (l + r) / 2

		res := f(m)

		if res > 0 {
			r = m
		} else {
			l = m
		}
	}

	fmt.Printf("%.6f\n", l)
}
