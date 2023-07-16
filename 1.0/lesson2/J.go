package main

import "fmt"

const (
	max = 4000.0
	min = 30.0
)

func mainJ() {
	var n int
	fmt.Scan(&n)

	l := min
	r := max

	var curr_ton float64
	fmt.Scan(&curr_ton)

	var cond string
	var new_ton float64
	for i := 1; i < n; i++ {
		fmt.Scan(&new_ton, &cond)

		if (cond == "closer" && new_ton > curr_ton) || (cond == "further" && curr_ton >= new_ton) {
			l, r = TryUpdateBorder(l, r, (new_ton+curr_ton)/2, max)
		} else {
			l, r = TryUpdateBorder(l, r, min, (new_ton+curr_ton)/2)
		}

		curr_ton = new_ton
	}

	fmt.Printf("%.8f %.8f\n", l, r)
}

func TryUpdateBorder(l, r, nl, nr float64) (float64, float64) {
	if nl >= l {
		l = nl
	}

	if nr <= r {
		r = nr
	}

	return l, r
}
