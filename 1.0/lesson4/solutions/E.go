package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainE() {
	sc := bufio.NewScanner(os.Stdin)
	var n int
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	dict := make(map[int]int)
	ans := 0

	for i := 0; i < n; i++ {
		var w, h int
		sc.Scan()
		input := sc.Text()
		arr := strings.Fields(input)
		w, _ = strconv.Atoi(arr[0])
		h, _ = strconv.Atoi(arr[1])

		if h > dict[w] {
			ans -= dict[w]
			ans += h
			dict[w] = h
		}
	}

	fmt.Println(ans)
}
