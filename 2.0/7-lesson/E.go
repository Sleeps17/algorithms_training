package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const pi float64 = math.Pi

type Event struct {
	Phi  float64
	Type int
}

func mainE() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	var rmin, rmax float64 = 0, 1e6 + 1

	events := make([]Event, 0, n)

	for i := 0; i < n; i++ {
		var r1, r2, phi1, phi2 float64
		fmt.Fscan(rd, &r1, &r2, &phi1, &phi2)

		rmin = max(rmin, r1)
		rmax = min(rmax, r2)

		events = append(events, Event{phi1, -(i + 1)})
		events = append(events, Event{phi2, (i + 1)})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Phi < events[j].Phi
	})

	used := make(map[int]bool)
	cntsegs := 0

	for _, ev := range events {
		if ev.Type < 0 {
			cntsegs++
			used[-ev.Type] = true
		} else if used[ev.Type] {
			cntsegs--
		}
	}

	ans := float64(0)

	for i := range events {
		ev := events[i]

		if ev.Type < 0 {
			cntsegs++
		} else {
			cntsegs--
		}

		if cntsegs == n {
			if i < len(events)-1 {
				ans += (events[i+1].Phi - events[i].Phi) * (rmax*rmax - rmin*rmin) / 2
				fmt.Println("1111", events[i+1].Phi, events[i].Phi, rmax, rmin)
			} else {
				ans += (events[0].Phi - events[len(events)-1].Phi + 2*pi) * (rmax*rmax - rmin*rmin) / 2
				fmt.Println("2222", events[0].Phi, events[len(events)-1].Phi, rmax, rmin)
			}
		}
	}
	fmt.Printf("%.6f\n", ans)
}
