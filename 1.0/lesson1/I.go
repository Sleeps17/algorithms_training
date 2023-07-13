package main

import "fmt"

func precalc(a ...int) (first int, second int) {
	if a[0] < a[1] {
		first = a[0]
		second = a[1]
	} else {
		first = a[1]
		second = a[0]
	}

	if len(a) == 2 {
		return
	}

	if a[2] < first {
		second = first
		first = a[2]
	} else if a[2] < second {
		second = a[2]
	}

	return
}

func mainI() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	f1, s1 := precalc(a, b, c)
	f2, s2 := precalc(d, e)

	if f2 >= f1 && s2 >= s1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
