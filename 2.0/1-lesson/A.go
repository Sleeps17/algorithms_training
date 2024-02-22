package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainA() {
	rd := bufio.NewReader(os.Stdin)

	var r, i, c int
	fmt.Fscan(rd, &r, &i, &c)

	ans := 0

	switch i {
	case 0:
		if r != 0 {
			ans = 3
		} else {
			ans = c
		}
	case 1:
		ans = c
	case 4:
		if r != 0 {
			ans = 3
		} else {
			ans = 4
		}
	case 6:
		ans = 0
	case 7:
		ans = 1
	default:
		ans = i
	}

	fmt.Println(ans)
}
