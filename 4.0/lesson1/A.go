package main

import (
	"bufio"
	"fmt"
	"os"
)

func partition(arr []int, pivot int) (int, int) {
	lowerCount := 0
	for _, num := range arr {
		if num < pivot {
			lowerCount++
		}
	}
	return lowerCount, len(arr) - lowerCount
}

func mainA() {
	rd := bufio.NewReader(os.Stdin)
	var n, pivot int
	fmt.Fscanf(rd, "%d\n", &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d", &arr[i])
	}
	fmt.Fscanf(rd, "\n%d", &pivot)

	lower, higher := partition(arr, pivot)
	fmt.Println(lower)
	fmt.Println(higher)
}
