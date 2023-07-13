package main

import (
	"fmt"
)

func presolutionJ(a, b, c, d, e, f float64) int {

	if a == 0 && b == 0 && c == 0 && d == 0 && e == 0 && f == 0 {
		return 5
	}
	if a == 0 && b == 0 && c == 0 && d == 0 {
		return 0
	}
	if b == 0 && d == 0 {
		if a != 0 && c != 0 {
			return 3
		}
	}
	if a == 0 && c == 0 {
		if b != 0 && d != 0 {
			return 4
		}
	}

	det := a*d - b*c
	det1 := e*d - b*f
	det2 := a*f - e*c

	if det == 0 {
		if det1 == 0 && det2 == 0 {
			if a == 0 && b == 0 && e == 0 {
				return 6
			}
			if c == 0 && d == 0 && f == 0 {
				return 7
			}
			return 1
		} else {
			return 0
		}
	} else {
		return 2
	}

}

func solutionJ(flag int, a, b, c, d, e, f float64) int {
	if flag == 0 {
		fmt.Println(0)
		return 0
	}
	if flag == 1 {
		fmt.Println(1, -a/b, e/b)
		return 0
	}
	if flag == 2 {

		det := a*d - b*c
		det1 := e*d - b*f
		det2 := a*f - e*c
		sol1 := det1 / det
		sol2 := det2 / det
		fmt.Println(2, sol1, sol2)
		return 0
	}
	if flag == 3 {

		x1 := e / a
		x2 := f / c
		if x1 == x2 {
			fmt.Println(3, x1)
		} else {
			fmt.Println(0)
		}
		return 0
	}
	if flag == 4 {

		y1, y2 := e/b, f/d
		if y1 == y2 {
			fmt.Println(4, y1)
		} else {
			fmt.Println(0)
		}
		return 0
	}
	if flag == 5 {

		fmt.Println(5)
		return 0
	}
	if flag == 6 {

		if d == 0 {
			fmt.Println(3, f/c)
			return 0
		}
		if c == 0 {
			fmt.Println(4, f/d)
			return 0
		}
		fmt.Println(1, -c/d, f/d)
		return 0
	}
	if flag == 7 {

		if b == 0 {
			fmt.Println(3, e/a)
			return 0
		}
		if a == 0 {
			fmt.Println(4, e/b)
			return 0
		}
		fmt.Println(1, -a/b, e/b)
		return 0
	}
	return 0

}

func mainJ() {
	var a, b, c, d, e, f float64
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	flag := presolutionJ(a, b, c, d, e, f)
	solutionJ(flag, a, b, c, d, e, f)

}
