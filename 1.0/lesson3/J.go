package main

import "fmt"

type point struct {
	x int
	y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func get_dist(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func make_naavigator_points(navigator_points map[point]bool, dx, dy [5]int, x, y, d int) {
	k := 0
	navigator_points[point{x, y}] = true
	tmp := navigator_points

	for k < d {
		for p, _ := range navigator_points {
			for j := 0; j < 5; j++ {
				tmp[point{p.x + dx[j], p.y + dy[j]}] = true
			}
		}
		navigator_points = tmp
		k++
	}
}

func mainJ() {
	dx := [5]int{0, 1, -1, 0, 0}
	dy := [5]int{1, 0, 0, -1, 0}

	var t, d, n int
	fmt.Scan(&t, &d, &n)

	starts := make(map[point]bool)
	starts[point{0, 0}] = true

	for i := 0; i < n; i++ {
		x, y := 0, 0
		fmt.Scan(&x, &y)

		navigator_points := make(map[point]bool)
		make_naavigator_points(navigator_points, dx, dy, x, y, d)

		mb_here := make(map[point]bool)

		for fp, _ := range navigator_points {
			for st, _ := range starts {
				if get_dist(fp, st) <= t {
					mb_here[fp] = true
					break
				}
			}
		}

		starts = mb_here
	}

	fmt.Println(len(starts))
	for p, _ := range starts {
		fmt.Println(p.x, p.y)
	}

}
