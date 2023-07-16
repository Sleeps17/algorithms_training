package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MakeUpdateG(number, max1, max2, min1, min2 *int) {
	if *number >= *max1 {
		*max2 = *max1
		*max1 = *number
	} else if *number > *max2 {
		*max2 = *number
	}

	if *number <= *min1 {
		*min2 = *min1
		*min1 = *number
	} else if *number < *min2 {
		*min2 = *number
	}

}

func mainG() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1000000), 1000000)
	sc.Scan()
	input := sc.Text()
	arr := strings.Split(input, " ")
	min1, _ := strconv.Atoi(arr[0])
	min2, _ := strconv.Atoi(arr[1])
	max1, _ := strconv.Atoi(arr[0])
	max2, _ := strconv.Atoi(arr[1])

	if min2 < min1 {
		tmp := min1
		min1 = min2
		min2 = tmp
	}

	if max1 < max2 {
		tmp := max1
		max1 = max2
		max2 = tmp
	}

	for i := 2; i < len(arr); i++ {
		number, _ := strconv.Atoi(arr[i])
		MakeUpdateG(&number, &max1, &max2, &min1, &min2)
	}

	if max1*max2 > min1*min2 {
		fmt.Println(max2, max1)
	} else {
		fmt.Println(min1, min2)
	}
}
