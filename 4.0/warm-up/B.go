package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func reduceFraction(a, b int) (int, int) {
	g := gcd(a, b)
	return a / g, b / g
}

func mainB() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	numerator := a*d + b*c
	denominator := b * d

	numerator, denominator = reduceFraction(numerator, denominator)

	fmt.Println(numerator, denominator)
}
