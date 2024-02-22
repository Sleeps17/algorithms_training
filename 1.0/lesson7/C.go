package main

import (
	"sort"
)

type stud struct {
	variant int64
	index   int64
}

func mainC() {
	defer in.Close()
	sc.Buffer(make([]byte, 10000000), 10000000)
	defer out.Close()

	n, d := ReadTwoInts()
	if n == 0 {
		return
	}
	events := ReadSliceEvents(d)
	studs := make(map[int64]stud)

	i := int64(0)
	for _, elem := range events {
		if elem.Type == -1 {
			studs[elem.Time+d] = stud{0, i}
			i++
		}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].Time == events[j].Time {
			return events[i].Type < events[j].Type
		}

		return events[i].Time < events[j].Time
	})

	cnt := int64(0)
	max := int64(0)
	varints := make([]int64, 0)

	for _, ev := range events {
		if ev.Type == -1 {
			cnt++
			if cnt > max {
				max = cnt
				a := studs[ev.Time+d]
				studs[ev.Time+d] = stud{cnt, a.index}

			} else {
				a := studs[ev.Time+d]
				studs[ev.Time+d] = stud{varints[0], a.index}
				varints = varints[1:]
			}
		} else {
			cnt--
			varints = append(varints, studs[ev.Time].variant)
		}
	}

	ans := make([]stud, 0, len(studs))
	for _, value := range studs {
		ans = append(ans, value)
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].index < ans[j].index
	})
	WriteInt(max, "\n")
	for i := 0; i < len(ans); i++ {
		if i == len(ans)-1 {
			WriteInt(ans[i].variant, "\n")
		} else {
			WriteInt(ans[i].variant, " ")
		}
	}
	wr.Flush()
}
