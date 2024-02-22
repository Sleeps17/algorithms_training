package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainB() {
	rd := bufio.NewReader(os.Stdin)

	var n, i, j int
	fmt.Fscan(rd, &n, &i, &j)

	if i == j {
		fmt.Println(0)
		return
	}

	if i > j {
		i, j = j, i
	}

	ans := min(j-i-1, (i+n-j)%n-1)
	fmt.Println(ans)
}
