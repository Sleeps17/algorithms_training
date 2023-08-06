package main

import (
	"bufio"
	"os"
	"sort"
)

func minB(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func mainB() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 100000000), 100000000)
	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	n, k := ReadTwoInts(sc)
	if n == 0 {
		return
	}
	arr1 := ReadSliceInts(sc)
	arr2 := ReadSliceInts(sc)

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})

	for i := int64(0); i < k; i++ {
		a := lower_bound(arr1, arr2[i])
		b := upper_bound(arr1, arr2[i])

		if a != b {
			WriteInt(wr, arr1[a], "\n")
		} else {
			if b == int64(len(arr1)) {
				WriteInt(wr, arr1[b-1], "\n")
			} else if a == 0 {
				WriteInt(wr, arr1[a], "\n")
			} else if arr2[i]-arr1[a-1] < arr1[b]-arr2[i] {
				WriteInt(wr, arr1[a-1], "\n")
			} else if arr2[i]-arr1[a-1] == arr1[b]-arr2[i] {
				WriteInt(wr, minB(arr1[a-1], arr1[b]), "\n")
			} else {
				WriteInt(wr, arr1[b], "\n")
			}
		}
	}
	wr.Flush()
}
