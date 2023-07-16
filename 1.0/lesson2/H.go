package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func kth_rearrange(arr []int, left, right, k int) {
	for left < right {
		x := arr[(left+right)/2]
		eqxfirst := left
		gtxfirst := left

		for i := left; i <= right; i++ {
			now := arr[i]

			if now == x {
				arr[i] = arr[gtxfirst]
				arr[gtxfirst] = now
				gtxfirst++
			} else if now < x {
				arr[i] = arr[gtxfirst]
				arr[gtxfirst] = x
				arr[eqxfirst] = now
				gtxfirst++
				eqxfirst++
			}
		}

		if k < eqxfirst {
			right = eqxfirst - 1
		} else if k >= eqxfirst {
			left = gtxfirst
		} else {
			return
		}
	}
}

func mainH() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	numbers := strings.Split(input, " ")
	var arr []int
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}
		arr = append(arr, num)
	}

	kth_rearrange(arr, 0, len(arr)-1, len(arr)-1)
	kth_rearrange(arr, 0, len(arr)-2, len(arr)-2)
	kth_rearrange(arr, 0, len(arr)-3, 2)

	if arr[len(arr)-1]*arr[len(arr)-2]*arr[len(arr)-3] > arr[len(arr)-1]*arr[0]*arr[1] {
		fmt.Println(arr[len(arr)-3], arr[len(arr)-1], arr[len(arr)-2])
	} else {
		fmt.Println(arr[len(arr)-1], arr[0], arr[1])
	}

}
