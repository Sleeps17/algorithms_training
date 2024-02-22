package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainB() {
	rd := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	set := make(map[int]bool)

	for {
		var num int

		if n, _ := fmt.Fscan(rd, &num); n != 1 {
			break
		}

		if set[num] {
			w.WriteString("YES\n")
		} else {
			w.WriteString("NO\n")
		}

		set[num] = true
	}
}
