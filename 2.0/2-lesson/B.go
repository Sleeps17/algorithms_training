package main

import (
	"fmt"
	"math"
)

func mainB() {
	var buildings [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&buildings[i])
	}

	maxDistance := 0
	for i := 0; i < 10; i++ {
		if buildings[i] == 1 {
			distance := findNearestShop(buildings, i)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	fmt.Println(maxDistance)
}

func findNearestShop(buildings [10]int, index int) int {
	distance := math.MaxInt32
	for i := 0; i < 10; i++ {
		if buildings[i] == 2 {
			currDistance := int(math.Abs(float64(index - i)))
			if currDistance < distance {
				distance = currDistance
			}
		}
	}
	return distance
}
