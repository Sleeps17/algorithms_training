package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func mainA() {
	rd := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscan(rd, &n)

	m := make(map[int64]int64)

	for i := int64(0); i < n; i++ {
		var d, a int64
		fmt.Fscan(rd, &d, &a)

		m[d] += a
	}

	var keys []int64
	for k := range m {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}
