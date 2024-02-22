package main

import (
	"fmt"
	"math"
)

func mainG() {
	var n, m int
	fmt.Scan(&n, &m)

	garden := make([][]int, n)
	for i := 0; i < n; i++ {
		garden[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&garden[i][j])
		}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	maxSide := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = garden[i][j]
			} else if garden[i][j] == 1 {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), math.Min(float64(dp[i][j-1]), float64(dp[i-1][j-1]))) + 1)
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}

	fmt.Println(maxSide)
}
