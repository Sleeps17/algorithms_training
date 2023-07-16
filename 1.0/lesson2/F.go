package main

import "fmt"

func isSymetric(arr []int, size, start int) bool {
	for i := 0; i < (size-start)/2; i++ {
		if arr[i+start] != arr[size-1-i] {
			return false
		}
	}

	return true
}

func mainF() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	ind := 0
	for i := 0; i < n; i++ {
		if isSymetric(arr, n, i) {
			ind = i
			break
		}
	}

	if ind == 0 {
		fmt.Println(ind)
	} else {
		fmt.Println(ind)
		for i := ind - 1; i >= 0; i-- {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()
}
