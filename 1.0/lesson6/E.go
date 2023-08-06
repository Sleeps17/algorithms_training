package main

import (
	"bufio"
	"os"
)

func mainE() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 10000000), 10000000)
	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	a := ReadInt(sc)
	b := ReadInt(sc)
	c := ReadInt(sc)

	l := int64(0)
	r := a + b + c
	for l < r {
		m := (l + r) / 2

		sum := 2*a + 3*b + 4*c + 5*m
		cnt := a + b + c + m

		if (sum/cnt < 3) || (sum/cnt == 3 && (sum%cnt < (cnt+1)/2)) {
			l = m + 1
		} else {
			r = m
		}
	}

	WriteInt(wr, l, "\n")
	wr.Flush()
}
