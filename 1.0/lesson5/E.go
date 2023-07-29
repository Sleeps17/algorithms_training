package main

import (
	"bufio"
	"os"
	"strconv"
)

func mainE() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 1000000000), 1000000000)
	n, k := read_two_ints(sc)

	arr := read_slice_ints(sc)

	sorts := make([]int, k+1)

	l := int64(0)
	r := int64(0)
	ansl, ansr := int64(0), int64(0)
	cnt := int64(0)

	for r < n && l < n {
		if r < n && cnt != k {
			if sorts[arr[r]] == 0 {
				cnt++
			}
			sorts[arr[r]]++
			r++
		} else {
			for cnt == k {
				sorts[arr[l]]--
				if sorts[arr[l]] == 0 {
					cnt--
				}
				l++
			}

			if r-l < ansr-ansl || (ansl == 0 && ansr == 0) {
				ansl = l
				ansr = r
			}
		}
	}

	if r == n && cnt == k {
		for cnt == k {
			sorts[arr[l]]--
			if sorts[arr[l]] == 0 {
				cnt--
			}
			l++
		}

		if r-l < ansr-ansl || (ansl == 0 && ansr == 0) {
			ansl = l
			ansr = r
		}
	}

	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)
	wr.Write([]byte(strconv.Itoa(int(ansl))))
	wr.Write([]byte(" "))
	wr.Write([]byte(strconv.Itoa(int(ansr))))
	wr.Write([]byte("\n"))
	wr.Flush()
}
