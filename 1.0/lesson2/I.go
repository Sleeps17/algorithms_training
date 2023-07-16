package main

import "fmt"

func MakeField(n, m int, minex, miney []int) {
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	field := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		field[i] = make([]int, m+2)
	}

	for i := 0; i < len(minex); i++ {
		x, y := minex[i], miney[i]

		for j := 0; j < len(dx); j++ {
			field[x+dx[j]][y+dy[j]] = field[x+dx[j]][y+dy[j]] + 1
		}
	}

	for i := 0; i < len(minex); i++ {
		field[minex[i]][miney[i]] = -1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if field[i][j] == -1 {
				fmt.Print("* ")
			} else {
				fmt.Printf("%d ", field[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func mainI() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	minex := make([]int, k)
	miney := make([]int, k)

	for i := 0; i < k; i++ {
		fmt.Scan(&minex[i], &miney[i])
	}

	MakeField(n, m, minex, miney)

}
