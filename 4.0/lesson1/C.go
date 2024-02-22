package main

import (
	"bufio"
	"os"
)

func merge(arr1, arr2 []int) []int {
	i := 0
	j := 0

	n := len(arr1)
	m := len(arr2)

	result := make([]int, 0, n+m)

	for i < n && j < m {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < n {
		result = append(result, arr1[i])
		i++
	}

	for j < m {
		result = append(result, arr2[j])
		j++
	}

	return result
}

func mainC() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 1000000)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	_ = ReadInt(sc)
	arr1 := ReadArr(sc)

	_ = ReadInt(sc)
	arr2 := ReadArr(sc)

	result := merge(arr1, arr2)
	WriteArr(w, result, " ", "\n")
}
