package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func mainG() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 100000000), 100000000)

	n, k := read_two_ints(sc)
	nums := read_slice_ints(sc)

	if n == 0 {
		return
	}

	cntnums := make(map[int64]int64)

	for _, num := range nums {
		cntnums[num]++
	}

	uniqnums := make([]int64, 0, len(cntnums))
	for k2 := range cntnums {
		uniqnums = append(uniqnums, k2)
	}
	sort.Slice(uniqnums, func(i, j int) bool {
		return uniqnums[i] < uniqnums[j]
	})
	fmt.Println(uniqnums)

	r := int64(0)
	ans := int64(0)
	dupl := int64(0)

	for l := int64(0); l < int64(len(uniqnums)); l++ {
		for r < int64(len(uniqnums)) && uniqnums[l]*k >= uniqnums[r] {
			if cntnums[uniqnums[r]] >= 2 {
				dupl++
			}
			r++
		}
		rangelen := r - l
		if cntnums[uniqnums[l]] >= 2 {
			ans += (rangelen - 1) * 3
		}
		if cntnums[uniqnums[l]] >= 3 {
			ans++
		}
		ans += (rangelen - 1) * (rangelen - 2) * 3
		if cntnums[uniqnums[l]] >= 2 {
			dupl--
		}
		ans += dupl * 3
	}

	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	write_int(wr, ans, "\n")
	wr.Flush()

}
