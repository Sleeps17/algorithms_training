package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainD() {
	rd := bufio.NewReader(os.Stdin)
	var s, i, n int
	var a [100001]int

	fmt.Fscan(rd, &n)
	s = n / 2
	for i = 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	if n%2 == 0 {
		fmt.Println(a[s])
	} else {
		fmt.Println(a[s+1])
	}
}
