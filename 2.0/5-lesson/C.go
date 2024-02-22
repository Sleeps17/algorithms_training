package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func mainC() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	arr1 := make(pairSlice, n)
	for i := range arr1 {
		var elem int64
		fmt.Fscan(rd, &elem)
		arr1[i] = pair{elem, int64(i)}
	}

	arr2 := make(pairSlice, m)
	for i := range arr2 {
		var elem int64
		fmt.Fscan(rd, &elem)
		arr2[i] = pair{elem, int64(i)}
	}

	sort.Sort(arr1)
	sort.Sort(arr2)

	f := 0
	s := 0

	cnt := 0
	ans := make([]int64, n)

	for f < n && s < m {
		if arr2[s].first > arr1[f].first {
			cnt++
			ans[arr1[f].second] = arr2[s].second + 1

			s++
			f++
		} else {
			s++
		}
	}

	fmt.Println(cnt)
	for i := range ans {
		fmt.Println(ans[i])
	}
}
