package main

import "fmt"

func mainE() {
	var n int
	fmt.Scan(&n)

	seq := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&seq[i])
	}

	discLow := make([]int, n)
	discHigh := make([]int, n)

	var spred = 0
	for i := 1; i < n; i++ {
		spred = seq[i] - seq[i-1]
		discLow[i] = discLow[i-1] + spred*i
	}

	for i := 0; i < n-1; i++ {
		discHigh[i] = (seq[n-1]-seq[i])*n - (discLow[n-1] - discLow[i])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", discLow[i]+discHigh[i])
	}
	fmt.Println()
}
