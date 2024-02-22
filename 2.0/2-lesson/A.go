package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainA() {
	rd := bufio.NewReader(os.Stdin)

	elem := -1
	max := -1
	cnt := 0

	for elem != 0 {
		fmt.Fscan(rd, &elem)

		if elem > max {
			max = elem
			cnt = 0
		}

		if elem == max {
			cnt++
		}
	}

	fmt.Println(cnt)
}
