package main

import (
	"sort"
)

type pairB struct {
	data  int64
	index int64
}

func MinB(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func MaxB(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func mainB() {
	defer in.Close()
	sc.Buffer(make([]byte, 10000000), 10000000)
	defer out.Close()

	n, m := ReadTwoInts()
	if m == 0 {
		return
	}
	events := make([]event, 2*n)
	for i := int64(0); i < 2*n; i += 2 {
		a, b := ReadTwoInts()
		events[i].Time, events[i+1].Time = MinB(a, b), MaxB(a, b)
		events[i].Type, events[i+1].Type = 1, 2
	}

	points := make([]pairB, 0, m)
	i := int64(0)
	for _, elem := range ReadSliceInts() {
		points = append(points, pairB{elem, i})
		i++
	}

	sort.Slice(events, func(i, j int) bool {

		if events[i].Time == events[j].Time {
			return events[i].Type < events[j].Type
		}

		return events[i].Time < events[j].Time
	})

	sort.Slice(points, func(i, j int) bool {
		return points[i].data < points[j].data
	})

	cnt := int64(0)
	i = int64(0)
	ans := make([]int64, m)

	for j := int64(0); j < 2*n; j++ {

		if events[j].Type == 1 {
			for i < m && points[i].data < events[j].Time {
				ans[points[i].index] = cnt
				i++
			}
			cnt++
		} else {
			for i < m && points[i].data <= events[j].Time {
				ans[points[i].index] = cnt
				i++
			}
			cnt--
		}
	}

	WriteSliceInts(ans, " ", "\n")
	wr.Flush()
}
