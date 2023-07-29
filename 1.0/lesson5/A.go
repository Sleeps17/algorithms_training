package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainA() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	n := read_int(sc)
	tshirts := read_slice_ints(sc)

	m := read_int(sc)
	pants := read_slice_ints(sc)

	i := int64(0)
	j := int64(0)
	ans := pair{tshirts[0], pants[0]}
	bestDiff := abs_A(tshirts[0] - pants[0])

	for i < n && j < m {
		diff := abs_A(tshirts[i] - pants[j])

		if diff < bestDiff {
			bestDiff = diff
			ans = pair{tshirts[i], pants[j]}
		}

		if tshirts[i] > pants[j] {
			j++
		} else {
			i++
		}
	}

	fmt.Println(ans.first, ans.second)

}

func abs_A(a int64) int64 {
	if a < 0 {
		return -a
	}

	return a
}
