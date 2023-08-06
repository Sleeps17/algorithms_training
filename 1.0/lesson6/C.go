package main

import (
	"bufio"
	"os"
)

func minC(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func maxC(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func mainC() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 1000000), 1000000)
	out, _ := os.Create("output.txt")
	defer out.Close()
	wrt := bufio.NewWriter(out)

	arr := ReadSliceInts(sc)
	w, h, n := arr[0], arr[1], arr[2]

	l := int64(1)
	r := maxC(w, h) * n

	for l < r {
		m := (l + r) / 2

		cntw := m / w
		cnth := m / h

		if cnth*cntw < n {
			l = m + 1
		} else {
			r = m
		}
	}

	WriteInt(wrt, l, "\n")
	wrt.Flush()

}
