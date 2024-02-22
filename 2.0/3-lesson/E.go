package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainE() {
	rd := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var m int
	fmt.Fscan(rd, &m)

	arr := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &arr[i])
	}

	var n int
	fmt.Fscan(rd, &n)

	maps := make([]map[rune]bool, n)
	nums := make([]string, n)
	set := make(map[string]bool)

	for i := 0; i < n; i++ {
		maps[i] = make(map[rune]bool)

		var str string
		fmt.Fscan(rd, &str)

		nums[i] = str

		for _, elem := range str {
			maps[i][elem] = true
		}
	}

	ans := make(map[string]int)

	for i := 0; i < n; i++ {
		if set[nums[i]] {
			continue
		}

		set[nums[i]] = true
		for j := 0; j < m; j++ {
			f := true
			for _, elem := range arr[j] {
				if maps[i][elem] == false {
					f = false
					break
				}

			}

			if f {
				ans[nums[i]]++
			}
		}
	}

	max := 0
	for _, value := range ans {
		if value > max {
			max = value
		}
	}

	for _, elem := range nums {
		if ans[elem] == max {
			w.WriteString(elem)
			w.WriteString("\n")
		}
	}
}
