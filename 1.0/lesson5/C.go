package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainC() {
	sc := bufio.NewScanner(os.Stdin)
	n := read_int(sc)
	arr := make([]int64, n)
	pref := make([]int64, n+1)
	rev_pref := make([]int64, n+1)

	prey := int64(0)

	for i := int64(0); i < n; i++ {
		_, y := read_two_ints(sc)

		arr[i] = y
	}

	prey = int64(0)
	for i := int64(1); i <= n; i++ {
		y := arr[i-1]
		if i == 1 {
			prey = y
		}

		if y-prey > 0 {
			pref[i] = pref[i-1] + (y - prey)
		} else {
			pref[i] = pref[i-1]
		}

		prey = y
	}

	for i := n - 1; i >= 0; i-- {
		y := arr[i]

		if i == n-1 {
			prey = y
		}

		if y-prey > 0 {
			rev_pref[i+1] = rev_pref[i+2] + (y - prey)
		} else {
			if i+2 <= n {
				rev_pref[i+1] = rev_pref[i+2]

			} else {
				rev_pref[i+1] = 0
			}
		}

		prey = y
	}

	m := read_int(sc)
	for i := int64(0); i < m; i++ {
		a, b := read_two_ints(sc)

		if b > a {
			fmt.Println(pref[b] - pref[a])
		} else {
			fmt.Println(rev_pref[b] - rev_pref[a])
		}
	}

}
