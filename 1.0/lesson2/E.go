package main

import "fmt"

func TryToGetAns(arr []int, size, pos int) (ans int) {
	for i := 0; i < size; i++ {
		if i == pos {
			continue
		}

		if arr[i] > arr[pos] {
			ans++
		}
	}

	return
}

func mainE() {
	var n int
	fmt.Scan(&n)

	max := 0
	ind := -1
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])

		if arr[i] > max {
			max = arr[i]
			ind = i
		}
	}

	ans_pos := 0
	preAns := -1

	for i := 1; i < n-1; i++ {
		if arr[i]%10 != 5 {
			continue
		}

		if ind >= i {
			continue
		}

		if arr[i+1] >= arr[i] {
			continue
		}

		if preAns == -1 {
			preAns = arr[i]
			ans_pos = i
		} else if arr[i] > preAns {
			preAns = arr[i]
			ans_pos = i
		}
	}

	if preAns == -1 {
		fmt.Println(0)
	} else {
		fmt.Println(TryToGetAns(arr, n, ans_pos) + 1)
	}

}
