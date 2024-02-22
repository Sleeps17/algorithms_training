package main

import (
	"sort"
)

type people struct {
	Time   int64
	Type   int64
	Number int64
}

func LEN(dict, anti map[int64]bool) int {

	return len(dict) - len(anti)
}

func mainD() {
	sc.Buffer(make([]byte, 10000000), 10000000)
	defer in.Close()
	defer out.Close()

	n := ReadInt()
	peoples := make([]people, 0)
	for i := int64(0); i < n; i++ {
		start, end := ReadTwoInts()
		if end-start < 5 {
			continue
		}
		peoples = append(peoples, people{start, 1, i})
		peoples = append(peoples, people{end - 5, 2, i})
	}

	sort.Slice(peoples, func(i, j int) bool {
		if peoples[i].Time == peoples[j].Time {
			return peoples[i].Type < peoples[j].Type
		}

		return peoples[i].Time < peoples[j].Time
	})

	if len(peoples) == 0 {
		WriteSliceInts([]int64{0, 10, 20}, " ", "\n")
	} else if len(peoples) == 2 {
		WriteSliceInts([]int64{1, peoples[0].Time, peoples[0].Time + 10}, " ", "\n")
	} else {
		bestans := 0
		firsbest, secondbest := 0, 0
		firststad := make(map[int64]bool)
		anti := make(map[int64]bool)
		for i := 0; i < len(peoples); i++ {
			event1 := peoples[i]
			if event1.Type == 1 {
				firststad[event1.Number] = true
				if LEN(firststad, anti) > bestans {
					bestans = LEN(firststad, anti)
					firsbest = int(event1.Time)
					secondbest = int(event1.Time) + 5
				}
			}
			secondadcnt := 0
			for j := i + 1; j < len(peoples); j++ {
				event2 := peoples[j]
				if event2.Type == 1 && firststad[event2.Number] == false {
					secondadcnt++
				}
				if event2.Time-5 >= event1.Time && LEN(firststad, anti)+secondadcnt > bestans {
					bestans = LEN(firststad, anti) + secondadcnt
					firsbest = int(event1.Time)
					secondbest = int(event2.Time)
				}
				if event2.Type == 2 && firststad[event2.Number] == false {
					secondadcnt--
				}
			}
			if event1.Type == 2 {
				firststad[event1.Number] = false
				anti[event1.Number] = true
			}
		}

		WriteSliceInts([]int64{int64(bestans), int64(firsbest), int64(secondbest)}, " ", "\n")
	}

	wr.Flush()
}
