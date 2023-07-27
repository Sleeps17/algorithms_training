package main

import "fmt"

func mainE() {
	keyb := make(map[int]bool)

	num := 0
	for i := 0; i < 3; i++ {
		fmt.Scan(&num)
		keyb[num] = true
	}

	var s string
	fmt.Scan(&s)
	number := []byte(s)
	set := make(map[byte]bool)

	ans := 0
	for _, elem := range number {
		if set[elem] == false {
			if keyb[int(elem-'0')] == false {
				ans++
			}
			set[elem] = true
		}
	}

	fmt.Println(ans)
}
