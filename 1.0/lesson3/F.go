package main

import "fmt"

func mainF() {
	var line1, line2 string
	fmt.Scan(&line1, &line2)

	set1 := make(map[string]int)
	set2 := make(map[string]int)

	arr1, arr2 := []byte(line1), []byte(line2)

	for i := 0; i < len(line1)-1; i++ {
		gen := []byte{arr1[i], arr1[i+1]}
		set1[string(gen)]++
	}

	for i := 0; i < len(line2)-1; i++ {
		gen := []byte{arr2[i], arr2[i+1]}
		set2[string(gen)] = 1
	}

	ans := 0

	for key, val := range set1 {
		ans += set2[key] * val
	}

	fmt.Println(ans)
}
