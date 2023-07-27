package main

import "fmt"

func mainH() {
	var n int
	fmt.Scan(&n)

	set := make(map[int]bool)

	for i := 0; i < n; i++ {
		x, y := 0, 0
		fmt.Scan(&x, &y)

		set[x] = true
	}

	fmt.Println(len(set))
}
