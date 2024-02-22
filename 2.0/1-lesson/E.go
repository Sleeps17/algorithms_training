package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	x int
	y int
}

func distance(a point, b point) float64 {
	return math.Sqrt(float64((b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y)))
}

func mainE() {
	rd := bufio.NewReader(os.Stdin)

	var d, x, y int
	fmt.Fscan(rd, &d, &x, &y)

	m := [3]point{
		{0, 0},
		{d, 0},
		{0, d},
	}

	if x >= 0 && x <= d && y >= 0 && y <= d-x {
		fmt.Println(0)
		return
	}

	ans := 0
	min := math.MaxFloat64
	for i, p := range m {
		dst := distance(p, point{x, y})
		if dst < min {
			min = dst
			ans = i + 1
		}
	}

	fmt.Println(ans)
}
