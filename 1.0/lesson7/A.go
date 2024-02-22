package main

import (
	"sort"
)

func mainA() {
	defer in.Close()
	sc.Buffer(make([]byte, 10000000), 10000000)
	defer out.Close()

	n, m := ReadTwoInts()
	if n == 0 {
		return
	}
	events := make([]event, 2*m)
	for i := int64(0); i < 2*m; i += 2 {
		events[i].Time, events[i+1].Time = ReadTwoInts()
		events[i].Type = 1
		events[i+1].Type = 2
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].Time == events[j].Time {
			return events[i].Type < events[j].Type
		}

		return events[i].Time < events[j].Time
	})

	cnt := int64(0)
	lastzero := int64(0)
	ans := int64(0)

	for _, ev := range events {

		if ev.Type == 1 {
			if cnt == 0 {
				ans += ev.Time - lastzero
			}
			cnt++
		}

		if ev.Type == 2 {
			cnt--
			if cnt == 0 {
				lastzero = ev.Time + 1
			}
		}
	}
	ans += n - lastzero

	WriteInt(ans, "\n")
	wr.Flush()

}
