package main

import "fmt"

func GetEntAndFloor(flatno, flatsoonfloor, floors int) (int, int) {
	floorsbefore := (flatno - 1) / flatsoonfloor
	ent := floorsbefore/floors + 1
	floor := floorsbefore%floors + 1
	return ent, floor
}

func check(k1, m, k2, p2, n2, flatsoonfloor int) (int, int) {
	ent2, floor2 := GetEntAndFloor(k2, flatsoonfloor, m)

	if ent2 == p2 && floor2 == n2 {
		return GetEntAndFloor(k1, flatsoonfloor, m)
	}
	return -1, -1
}

func mainE() {
	var k1, m, k2, p2, n2 int
	fmt.Scan(&k1, &m, &k2, &p2, &n2)

	ent := -1
	floor := -1
	goodflag := false

	for flatsoonfloor := 1; flatsoonfloor <= int(1e6); flatsoonfloor++ {
		nent, nfloor := check(k1, m, k2, p2, n2, flatsoonfloor)

		if nent != -1 {
			goodflag = true
			if ent == -1 {
				ent, floor = nent, nfloor
			} else if ent != nent && ent != 0 {
				ent = 0
			} else if floor != nfloor && floor != 0 {
				floor = 0
			}
		}
	}

	if goodflag {
		fmt.Println(ent, floor)
	} else {
		fmt.Println(-1, -1)
	}
}
