package main

import (
	"bufio"
	"math"
	"os"
)

func mainH() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 1000000), 1000000)
	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	n, k := ReadTwoInts(sc)
	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		arr[i] = ReadInt(sc)
	}

	l := int64(0)
	r := int64(math.MaxInt64)
	for l < r {
		m := (l + r) / 2

		cnt := int64(0)
		for _, elem := range arr {
			if m == 0 {
				break
			}
			cnt += elem / m
		}

		if cnt >= k {
			l = m + 1
		} else {
			r = m
		}
	}

	if l == 0 {
		WriteInt(wr, 0, "\n")
	} else {
		WriteInt(wr, l-1, "\n")
	}
	wr.Flush()
}
