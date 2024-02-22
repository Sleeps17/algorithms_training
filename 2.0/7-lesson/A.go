package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type event struct {
	Coord  int64
	Action int64
}

type slice_event []event

func (s slice_event) Len() int {
	return len(s)
}

func (s slice_event) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s slice_event) Less(i, j int) bool {
	if s[i].Coord == s[j].Coord {
		return s[i].Action > s[j].Action
	}

	return s[i].Coord < s[j].Coord
}

func mainA() {
	rd := bufio.NewReader(os.Stdin)

	var n int64
	fmt.Fscan(rd, &n)

	events := make(slice_event, 0, 2*n)
	for i := int64(0); i < n; i++ {
		var l, r int64
		fmt.Fscan(rd, &l, &r)

		events = append(events, event{l, 1})
		events = append(events, event{r, -1})
	}

	sort.Sort(events)

	fmt.Println(events)

	cnt := int64(0)
	ans := int64(0)
	start := int64(0)

	for _, ev := range events {
		if ev.Action == 1 {
			if cnt == 0 {
				start = ev.Coord
			}
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				ans += ev.Coord - start
			}
		}
	}
	fmt.Println(ans)
}
