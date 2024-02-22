package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadSlice(sc *bufio.Scanner) []int {

	sc.Scan()
	strNums := strings.Fields(sc.Text())

	arr := make([]int, 0, len(strNums))

	for _, strNum := range strNums {
		num, _ := strconv.Atoi(strNum)

		arr = append(arr, num)
	}

	return arr
}

func mainA() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	arr1 := ReadSlice(sc)

	arr2 := ReadSlice(sc)

	set := make(map[int]int)

	for i := range arr1 {
		set[arr1[i]]++
	}

	ans := 0
	for i := range arr2 {
		if set[arr2[i]] > 0 {
			set[arr2[i]]--
			ans++
		}
	}

	fmt.Println(ans)
}
