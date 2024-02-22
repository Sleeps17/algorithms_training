package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type cargo struct {
	Time   int
	Action int
}

type slice_cargo []cargo

func (s slice_cargo) Len() int {
	return len(s)
}

func (s slice_cargo) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s slice_cargo) Less(i, j int) bool {
	if s[i].Time == s[j].Time {
		return s[i].Action < s[j].Action
	}

	return s[i].Time < s[j].Time
}

func mainB() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	cargos := make(slice_cargo, 0, 2*n)
	for i := 0; i < n; i++ {
		var t, l int
		fmt.Fscan(rd, &t, &l)

		cargos = append(cargos, cargo{t, 1})
		cargos = append(cargos, cargo{t + l, -1})
	}

	sort.Sort(cargos)

	cnt := 0
	ans := 0

	for _, c := range cargos {
		cnt += c.Action
		ans = max(cnt, ans)
	}

	fmt.Println(ans)
}
