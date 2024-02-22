package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainC() {
	rd := bufio.NewReader(os.Stdin)

	var x, y, z int
	fmt.Fscan(rd, &x, &y, &z)

	if x == y {
		fmt.Println(1)
		return
	}

	if x >= 1 && x <= 12 && y >= 1 && y <= 12 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
