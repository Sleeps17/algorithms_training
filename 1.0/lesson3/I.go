package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mainI() {
	n := 0
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)

	allknow := []string{}
	oneknow := []string{}
	set := make(map[string]int)

	for i := 0; i < n; i++ {
		m := 0
		sc.Scan()
		m, _ = strconv.Atoi(sc.Text())
		var lang string
		for j := 0; j < m; j++ {
			sc.Scan()
			lang = sc.Text()
			set[lang]++
		}
	}

	for lang, cnt := range set {
		if cnt == n {
			allknow = append(allknow, lang)
		}
		oneknow = append(oneknow, lang)
	}

	fmt.Println(len(allknow))
	for _, elem := range allknow {
		fmt.Println(elem)
	}
	fmt.Println(len(oneknow))
	for _, elem := range oneknow {
		fmt.Println(elem)
	}
}
