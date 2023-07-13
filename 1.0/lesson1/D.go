package main

import "fmt"

func mainD() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c < 0 {
		fmt.Println("NO SOLUTION")
	} else if b == c*c && a == 0 {
		fmt.Println("MANY SOLUTIONS")
	} else if a == 0 {
		fmt.Println("NO SOLUTION")
	} else if (c*c-b)%a == 0 {
		fmt.Println((c*c - b) / a)
	} else {
		fmt.Println("NO SOLUTION")
	}
}
