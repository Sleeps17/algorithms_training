package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func check(length int64, coords []int64, k int64) bool {
	begin := coords[0]
	var count int64 = 1
	for i := 1; i < len(coords); i++ {
		dist := coords[i] - begin
		if dist > length {
			count++
			begin = coords[i]
		}
	}
	return count <= k
}

func mainE() {
	rd := bufio.NewReader(os.Stdin)

	var n, k int64
	fmt.Fscan(rd, &n, &k)

	arr := make([]int64, n)
	for i := range arr {
		fmt.Fscan(rd, &arr[i])
	}

	slices.Sort(arr)

	var l, r int64 = -1, 1e11

	for r-l > 1 {
		mid := (l + r) / 2

		if check(mid, arr, k) {
			r = mid
		} else {
			l = mid
		}
	}

	fmt.Println(r)
}
