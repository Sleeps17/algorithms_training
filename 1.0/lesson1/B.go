package main

import "fmt"

func mainB() {
	a, b, c := 0, 0, 0
	fmt.Scan(&a, &b, &c)

	if a+b > c && a+c > b && b+c > a {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
