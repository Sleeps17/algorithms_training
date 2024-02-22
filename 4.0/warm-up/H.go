package main

import "fmt"

func mainH() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	cnt1 := a / 1
	cnt2 := (b + n - 1) / n

	if cnt1 > cnt2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
