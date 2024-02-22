package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainC() {
	rd := bufio.NewReader(os.Stdin)

	arr := make([]int, 0)
	set := make(map[int]int)

	for {

		var num int
		if n, _ := fmt.Fscan(rd, &num); n != 1 {
			break
		}

		if set[num] == 0 {
			arr = append(arr, num)
		}

		set[num]++
	}

	for _, elem := range arr {
		if set[elem] == 1 {
			fmt.Print(elem, " ")
		}
	}
	fmt.Println()
}
