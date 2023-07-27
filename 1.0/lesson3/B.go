package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func mainB() {
	rd := bufio.NewReader(os.Stdin)
	input1, _ := rd.ReadString('\n')
	input1 = strings.TrimSuffix(input1, "\n")
	arr1 := strings.Split(input1, " ")

	input2, _ := rd.ReadString('\n')
	input2 = strings.TrimSuffix(input2, "\n")
	arr2 := strings.Split(input2, " ")

	s1, s2 := NewSet(10), NewSet(10)

	for _, elem := range arr1 {
		num, _ := strconv.Atoi(elem)
		s1.add(num)
	}

	for _, elem := range arr2 {
		num, _ := strconv.Atoi(elem)
		s2.add(num)
	}

	res := make([]int, 0)

	for _, list := range s1.data {
		for _, num := range list {
			if s2.find(num) {
				res = append(res, num)
			}
		}
	}
	sort.Sort(sort.IntSlice(res))
	for _, elem := range res {
		fmt.Printf("%d ", elem)
	}
}
