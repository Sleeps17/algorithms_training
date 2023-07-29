package main

import (
	"bufio"
	"os"
)

func mainH() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 100000000), 100000000)

	n, k := read_two_ints(sc)
	str := []byte(read_string(sc))

	alph := make([]int64, 26)

	l := int64(0)
	r := int64(0)

	len_ans := int64(0)
	pos_ans := int64(0)

	for l < n && r < n {
		alph[str[r]-'a']++

		if alph[str[r]-'a'] == k+1 {

			if r-l > len_ans {
				len_ans = r - l
				pos_ans = l + 1
			}

			for alph[str[r]-'a'] == k+1 && l < n {
				alph[str[l]-'a']--
				l++
			}
		}

		r++
	}

	f := true

	if r == n {
		ll := l
		for ll < n {
			if alph[str[ll]-'a'] > k {
				f = false
				break
			}
			ll++
		}
	}

	if f == true {
		if r-l > len_ans {
			len_ans = r - l
			pos_ans = l + 1
		}
	}

	if len_ans == 0 && pos_ans == 0 {
		len_ans = int64(len(string(str)))
		pos_ans = 1
	}

	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	write_two_ints(wr, len_ans, pos_ans, " ", "\n")
	wr.Flush()
}
