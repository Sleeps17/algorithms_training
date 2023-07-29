package main

import (
	"bufio"
	"os"
)

func mainI() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 100000000), 100000000)

	k := read_int(sc)
	str := []byte(read_string(sc))
	n := int64(len(str))

	i := k
	ans := int64(0)
	prev := int64(0)

	for i < n {
		if str[i] == str[i-k] {
			prev++
			ans += prev
		} else {
			prev = 0
		}
		i++
	}

	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	write_int(wr, ans, "\n")
	wr.Flush()
}
