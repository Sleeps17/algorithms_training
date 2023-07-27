package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type buys map[string]int

func mainF() {
	shop := make(map[string]buys)

	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	for sc.Scan() {
		input := sc.Text()
		arr := strings.Fields(input)
		name := arr[0]
		item := arr[1]
		sum, _ := strconv.Atoi(arr[2])
		_, have := shop[name]
		if have == false {
			shop[name] = make(buys)
		}
		shop[name][item] += sum
	}

	sort_mans := make([]string, len(shop))

	i := 0
	for man, _ := range shop {
		sort_mans[i] = man
		i++
	}

	sort.Slice(sort_mans, func(i, j int) bool {
		return sort_mans[i] < sort_mans[j]
	})

	for _, man := range sort_mans {
		items := make([]string, len(shop[man]))
		i := 0
		for item, _ := range shop[man] {
			items[i] = item
			i++
		}
		sort.Slice(items, func(i, j int) bool {
			return items[i] < items[j]
		})

		wr.Write([]byte(man))
		wr.Write([]byte(":\n"))
		for _, item := range items {
			wr.Write([]byte(item))
			wr.Write([]byte(" "))
			int_sum, _ := shop[man][item]
			string_sum := strconv.Itoa(int_sum)
			wr.Write([]byte(string_sum))
			wr.Write([]byte("\n"))
		}
	}
	wr.Flush()
}
