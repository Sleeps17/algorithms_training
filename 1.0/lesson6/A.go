package main

import (
	"bufio"
	"os"
	"sort"
)

func mainA() {
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

		if a == b {
			WriteString(wr, "NO", "\n")
		} else {
			WriteString(wr, "YES", "\n")
		}
	}

	wr.Flush()
}
