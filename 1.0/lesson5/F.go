package main

import (
	"bufio"
	"os"
	"sort"
)

type pair struct {
	first  int64
	second int64
}

func mainF() {
	in, _ := os.Open("input.txt")
	defer in.Close()

	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 10000000), 10000000)
	n := read_int(sc)
	minPower := read_slice_ints(sc)

	m := read_int(sc)
	power := make([]pair, m)
	for i := int64(0); i < m; i++ {
		power[i].first, power[i].second = read_two_ints(sc)
	}

	sort.Slice(power, func(i, j int) bool {
		return power[i].second < power[j].second
	})
	sort.Slice(minPower, func(i, j int) bool {
		return minPower[i] < minPower[j]
	})

	i := int64(0)
	j := int64(0)
	ans := int64(0)

	for i < n && j < m {
		if minPower[i] <= power[j].first {
			ans += power[j].second
			i++
		} else {
			j++
		}
	}

	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)
	write_int(wr, ans, "\n")
	wr.Flush()

}
