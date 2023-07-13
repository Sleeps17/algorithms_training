package main

import (
	"fmt"
)

type elem struct {
	s int
	a int
	b int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mainF() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	l := make([]elem, 0, 4)
	l = append(l, elem{max(a, c) * (b + d), max(a, c), b + d})
	l = append(l, elem{max(a, d) * (b + c), max(a, d), b + c})
	l = append(l, elem{max(b, c) * (a + d), max(b, c), a + d})
	l = append(l, elem{max(b, d) * (a + c), max(b, d), a + c})

	min := l[0].s
	ind := 0
	for i, item := range l {
		if item.s < min {
			min = item.s
			ind = i
		}
	}

	fmt.Println(l[ind].a, l[ind].b)
}
