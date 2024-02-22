package main

import (
	"bufio"
	"fmt"
	"os"
)

func lower_bound(arr []int, goal int) int {
	var l, r = 0, len(arr)

	for l < r {
		mid := (l + r) / 2

		if arr[mid] < goal {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func upper_bound(arr []int, goal int) int {
	var l, r = 0, len(arr)

	for l < r {
		mid := (l + r) / 2

		if arr[mid] <= goal {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func mainB() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	arr1 := make([]int, n)
	for i := range arr1 {
		fmt.Fscan(rd, &arr1[i])
	}

	var m int
	fmt.Fscan(rd, &m)

	arr2 := make([]int, m)
	for i := range arr2 {
		fmt.Fscan(rd, &arr2[i])
	}

	for _, goal := range arr2 {
		f := lower_bound(arr1, goal)
		l := upper_bound(arr1, goal)

		if f == l {
			fmt.Println(0, 0)
		} else {
			fmt.Println(f+1, l)
		}
	}
}
