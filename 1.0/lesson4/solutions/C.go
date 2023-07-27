package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainC() {
	sc := bufio.NewScanner(os.Stdin)

	dict := make(map[string]int)
	arr := make([]string, 0)

	for sc.Scan() {
		arr = append(arr, strings.Fields(sc.Text())...)
	}

	for _, elem := range arr {
		dict[elem]++
	}

	max := 0
	word := ""
	for w, cnt := range dict {
		if cnt > max {
			word = w
			max = cnt
		} else if cnt == max && strings.Compare(w, word) == -1 {
			word = w
		}
	}

	fmt.Println(word)
}
