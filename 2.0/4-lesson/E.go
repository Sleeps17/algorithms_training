package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ReadInt(sc *bufio.Scanner) int {

	sc.Scan()

	num, _ := strconv.Atoi(sc.Text())

	return num
}

func mainE() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	n := ReadInt(sc)

	m := make(map[string]int)
	mp := make(map[string]int)
	topics := make(map[int]string)

	for i := 1; i <= n; i++ {
		k := ReadInt(sc)

		if k == 0 {
			var topic, msg string
			if sc.Scan() {
				topic = sc.Text()
			}

			m[topic]++
			mp[topic] = i
			topics[i] = topic

			if sc.Scan() {
				msg = sc.Text()
			}

			_ = msg
		} else {
			m[topics[k]]++
			topics[i] = topics[k]
			sc.Scan()
		}
	}

	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	res := make([]string, 0, n)

	for k, v := range m {
		if v == max {
			res = append(res, k)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return mp[res[i]] < mp[res[j]]
	})

	fmt.Println(res[0])
}
