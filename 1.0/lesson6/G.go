package main

import (
	"bufio"
	"math"
	"os"
)

func MaxG(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func MinG(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func mainG() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 10000000), 10000000)
	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	n := ReadInt(sc)
	m := ReadInt(sc)
	t := ReadInt(sc)

	l := int64(0)
	r := int64(math.MaxInt64)
	for l < r {
		mid := (l + r) / 2

		cnt := n*m - (n-2*mid)*(m-2*mid)

		if mid <= (MinG(n, m))/2 && cnt > 0 && cnt <= t {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if r == 0 {
		WriteInt(wr, 0, "\n")
	} else {
		WriteInt(wr, r-1, "\n")
	}
	wr.Flush()

}
