package main

import (
	"bufio"
	"os"
)

func MaxD(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func CalcD(wide, high, w, h, n int64) bool {
	cntw := w / wide
	cnth := h / high

	if cnth*cntw >= n {
		return true
	}

	cntw = w / high
	cnth = h / wide

	if cnth*cntw >= n {
		return true
	}

	return false
}

func mainD() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 1000000), 1000000)
	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	arr := ReadSliceInts(sc)
	n, a, b, w, h := arr[0], arr[1], arr[2], arr[3], arr[4]

	l := int64(0)
	r := MaxD(w, h)
	for l < r {
		m := (l + r) / 2

		if CalcD((a + 2*m), (b + 2*m), w, h, n) {
			l = m + 1
		} else {
			r = m
		}

	}

	WriteInt(wr, r-1, "\n")
	wr.Flush()
}
