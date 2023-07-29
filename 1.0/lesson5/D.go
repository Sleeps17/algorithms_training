package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainD() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)
	n, r := read_two_ints(sc)

	dist := read_slice_ints(sc)

	j := int64(0)
	i := int64(0)
	ans := int64(0)

	for i < n {
		if j < n && dist[j]-dist[i] > r {
			ans += (n - j)
			i++
		} else {
			if j == n-1 {
				break
			}
			j++
		}
	}

	fmt.Println(ans)
}
