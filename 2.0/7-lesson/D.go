package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	in = iota - 1
	cat
	out
)

type Point struct {
	Coord int
	Type  int
	Index int
}

func mainD() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	points := make([]Point, 0, n+2*m)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(rd, &a)

		points = append(points, Point{a, cat, -1})
	}
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Fscan(rd, &l, &r)

		points = append(points, Point{l, in, i})
		points = append(points, Point{r, out, i})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].Coord == points[j].Coord {
			return points[i].Type < points[j].Type
		}

		return points[i].Coord < points[j].Coord
	})

	ans := make([]int, m)
	cntCats := 0

	for _, point := range points {
		if point.Type == cat {
			cntCats++
		} else if point.Type == in {
			ans[point.Index] -= cntCats
		} else {
			ans[point.Index] += cntCats
		}
	}

	for _, elem := range ans {
		fmt.Println(elem)
	}
}
