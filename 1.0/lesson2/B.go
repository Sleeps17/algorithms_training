package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mainB() {
	sc := bufio.NewScanner(os.Stdin)
	cond := make([]bool, 6)
	for i := 0; i < len(cond); i++ {
		cond[i] = true
	}
	sc.Scan()
	prev, _ := strconv.Atoi(sc.Text())
	for sc.Scan() {
		number, _ := strconv.Atoi(sc.Text())
		if number == -2000000000 {
			break
		}

		if cond[0] != false && number != prev {
			cond[0] = false
		}
		if cond[1] != false && number <= prev {
			cond[1] = false
		}
		if cond[2] != false && number < prev {
			cond[2] = false
		}
		if cond[3] != false && number >= prev {
			cond[3] = false
		}
		if cond[4] != false && number > prev {
			cond[4] = false
		}

		prev = number
	}

	for ind, elem := range cond {
		if elem != true {
			continue
		}
		if ind == 0 {
			fmt.Println("CONSTANT")
			break
		} else if ind == 1 {
			fmt.Println("ASCENDING")
			break
		} else if ind == 2 {
			fmt.Println("WEAKLY ASCENDING")
			break
		} else if ind == 3 {
			fmt.Println("DESCENDING")
			break
		} else if ind == 4 {
			fmt.Println("WEAKLY DESCENDING")
			break
		} else if ind == 5 {
			fmt.Println("RANDOM")
			break
		}
	}

}
