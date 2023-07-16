package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainD() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	arrStr := strings.Split(input, " ")

	arr := make([]int, len(arrStr))
	for i := 0; i < len(arr); i++ {
		arr[i], _ = strconv.Atoi(arrStr[i])
	}

	cnt := 0
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			cnt++
		}
	}

	fmt.Println(cnt)

}
