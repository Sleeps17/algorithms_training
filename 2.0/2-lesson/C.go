package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainC() {
	sc := bufio.NewScanner(os.Stdin)

	var str string

	if sc.Scan() {
		str = sc.Text()
	}

	l := 0
	r := len(str) - 1

	cnt := 0

	for l < r {
		if str[l] != str[r] {
			cnt++
		}

		l++
		r--
	}

	fmt.Println(cnt)
}
