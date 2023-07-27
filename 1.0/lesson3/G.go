package main

import "fmt"

func mainG() {
	var n int
	fmt.Scan(&n)

	set := make(map[int]bool)

	for i := 0; i < n; i++ {
		a, b := 0, 0
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 {
			continue
		}
		if a+b+1 == n {
			set[a] = true
		}
	}

	fmt.Println(len(set))
}
