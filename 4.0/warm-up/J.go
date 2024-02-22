package main

import (
	"fmt"
)

func calculateResult() {
	var total, divisorA, divisorB int
	fmt.Scanf("%d %d %d", &total, &divisorA, &divisorB)
	minPerEach := total / divisorA
	remainder := total - total/divisorA*divisorA
	if remainder%minPerEach == 0 {
		if remainder/minPerEach <= divisorB-divisorA {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		if remainder/minPerEach+1 <= divisorB-divisorA {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func mainJ() {
	var testCases int
	fmt.Scanf("%d", &testCases)
	for i := 0; i < testCases; i++ {
		calculateResult()
	}
}
