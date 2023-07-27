package main

import (
	"fmt"
	"sort"
)

func mainC() {
	var n, m int
	fmt.Scan(&n, &m)

	s1 := make(map[int]bool)
	s2 := make(map[int]bool)

	for i := 0; i < n; i++ {
		num := 0
		fmt.Scan(&num)
		s1[num] = true
	}

	for i := 0; i < m; i++ {
		num := 0
		fmt.Scan(&num)
		s2[num] = true
	}

	res1, res2, res3 := make([]int, 0), make([]int, 0), make([]int, 0)

	for key, _ := range s1 {
		if s2[key] {
			res1 = append(res1, key)
		} else {
			res2 = append(res2, key)
		}
	}

	for key, _ := range s2 {
		if !s1[key] {
			res3 = append(res3, key)
		}
	}

	sort.Sort(sort.IntSlice(res1))
	sort.Sort(sort.IntSlice(res2))
	sort.Sort(sort.IntSlice(res3))

	fmt.Println(len(res1))
	for _, elem := range res1 {
		fmt.Printf("%d ", elem)
	}
	fmt.Println()
	fmt.Println(len(res2))
	for _, elem := range res2 {
		fmt.Printf("%d ", elem)
	}
	fmt.Println()
	fmt.Println(len(res3))
	for _, elem := range res3 {
		fmt.Printf("%d ", elem)
	}
	fmt.Println()
}
