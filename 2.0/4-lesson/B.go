package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func mainB() {
	rd := bufio.NewReader(os.Stdin)

	m := make(map[string]int64)

	for {

		var name string
		var cnt int64

		if n, _ := fmt.Fscan(rd, &name, &cnt); n != 2 {
			break
		}

		m[name] += cnt
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}
