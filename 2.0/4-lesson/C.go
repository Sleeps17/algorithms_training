package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	Key   int64
	Value string
}

type pairSlice []pair

func (s pairSlice) Len() int {
	return len(s)
}

func (s pairSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s pairSlice) Less(i, j int) bool {
	if s[i].Key == s[j].Key {
		return s[i].Value < s[j].Value
	}

	return s[i].Key > s[j].Key
}

func mainC() {
	rd := bufio.NewReader(os.Stdin)

	m := make(map[string]int64)

	for {
		var word string

		if n, _ := fmt.Fscan(rd, &word); n != 1 {
			break
		}

		m[word]++
	}

	var keys pairSlice
	for k, v := range m {
		keys = append(keys, pair{v, k})
	}

	sort.Sort(keys)

	for _, key := range keys {
		fmt.Println(key.Value)
	}
}
