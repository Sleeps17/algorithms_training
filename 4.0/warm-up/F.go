package main

import (
	"bufio"
	"fmt"
	"os"
)

func calculateTotalTime(maxCapacity int, peopleOnFloors []int) int64 {
	totalTime := int64(0)
	peopleRemainder := 0
	upFloor := 0
	for i := len(peopleOnFloors) - 1; i >= 0; i-- {
		floor := i + 1
		if peopleRemainder == 0 {
			upFloor = floor
		}
		peopleRemainder += peopleOnFloors[i]
		if peopleRemainder >= maxCapacity {
			totalTime += int64(upFloor * 2)
			peopleRemainder -= maxCapacity
			upFloor = floor
		}
		if peopleRemainder >= maxCapacity {
			trips := (peopleRemainder) / maxCapacity
			totalTime += int64(floor * 2 * trips)
			peopleRemainder %= maxCapacity
		}
	}
	if peopleRemainder != 0 {
		totalTime += int64(upFloor * 2)
	}
	return totalTime
}

func mainF() {

	rd := bufio.NewReader(os.Stdin)
	//start := time.Now()
	var maxCapacity, numFloors int
	fmt.Fscan(rd, &maxCapacity)
	fmt.Fscan(rd, &numFloors)
	peopleOnFloors := make([]int, numFloors)
	for i := 0; i < numFloors; i++ {
		fmt.Fscan(rd, &peopleOnFloors[i])
	}
	fmt.Println(calculateTotalTime(maxCapacity, peopleOnFloors))
	//fmt.Println(time.Since(start))
}
