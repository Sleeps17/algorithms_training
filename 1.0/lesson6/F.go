package main

import (
	"bufio"
	"math"
	"os"
)

func MinF(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func mainF() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 10000000), 10000000)
	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	arr := ReadSliceInts(sc)
	n, x, y := arr[0], arr[1], arr[2]

	l := int64(0)
	r := int64(math.MaxInt64)
	ans := MinF(x, y)
	for l < r {
		m := (l + r) / 2

		f := m/x + m/y

		if f < n-1 {
			l = m + 1
		} else {
			r = m
		}
	}

	WriteInt(wr, ans+r, "\n")
	wr.Flush()
}
