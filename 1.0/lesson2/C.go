package main

import "fmt"

func absC(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func mainC() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var x int
	fmt.Scan(&x)
	delta := absC(x - arr[0])
	ans := arr[0]

	for i := 0; i < n; i++ {
		if absC(x-arr[i]) < delta {
			delta = absC(x - arr[i])
			ans = arr[i]
		}
	}

	fmt.Println(ans)
}
