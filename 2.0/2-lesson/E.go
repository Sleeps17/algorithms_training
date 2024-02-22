package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainE() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	sup := -1
	sum := 0

	for i := 0; i < n; i++ {
		var elem int
		fmt.Fscan(rd, &elem)

		sum += elem
		sup = max(elem, sup)
	}

	fmt.Println(sum - sup)
}
