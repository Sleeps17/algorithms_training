package main

import (
	"bufio"
	"os"
)

func lomutoPartition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quicksortLomuto(arr []int, low int, high int) {
	if low < high {
		p := lomutoPartition(arr, low, high)
		quicksortLomuto(arr, low, p-1)
		quicksortLomuto(arr, p+1, high)
	}
}

func mainB() {
	sc := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	sc.Buffer(make([]byte, 10000000000), 10000000000)
	n := ReadInt(sc)
	arr := ReadArr(sc)

	quicksortLomuto(arr, 0, n-1)
	WriteArr(w, arr, " ", "\n")
	w.Flush()
}
