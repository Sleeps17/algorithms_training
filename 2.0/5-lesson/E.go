package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type triple struct {
	first  int64
	second int64
	third  int64
}

func Less(a, b triple) bool {
	if a.first < b.first {
		return true
	} else if a.first == b.first && a.second < b.second {
		return true
	} else if a.first == b.first && a.second == b.second && a.third < b.third {
		return true
	}

	return false
}

func Min(a, b triple) triple {
	if Less(a, b) {
		return a
	}

	return b
}

func TwoSum(a, b pairSlice, goal int64) (triple, bool) {
	f := 0
	s := len(b) - 1

	finded := false
	var ans = triple{math.MaxInt64, math.MaxInt64, math.MaxInt64}

	for f < len(a) && s >= 0 {

		if a[f].first+b[s].first == goal {
			// fmt.Println(a[f].first, b[s].first)
			ans = Min(ans, triple{a[f].second, b[s].second, 0})
			finded = true
		}

		if a[f].first+b[s].first <= goal {
			f++
		} else {
			s--
		}
	}

	//fmt.Println(ans, finded)

	return ans, finded
}

func mainE() {
	rd := bufio.NewReader(os.Stdin)

	var s int64
	fmt.Fscan(rd, &s)

	var sizeA int64
	fmt.Fscan(rd, &sizeA)
	a := make(pairSlice, sizeA)
	for i := range a {
		var elem int64
		fmt.Fscan(rd, &elem)

		a[i] = pair{elem, int64(i)}
	}

	var sizeB int64
	fmt.Fscan(rd, &sizeB)
	b := make(pairSlice, sizeB)
	for i := range b {
		var elem int64
		fmt.Fscan(rd, &elem)

		b[i] = pair{elem, int64(i)}
	}

	var sizeC int64
	fmt.Fscan(rd, &sizeC)
	c := make(pairSlice, sizeC)
	for i := range c {
		var elem int64
		fmt.Fscan(rd, &elem)

		c[i] = pair{elem, int64(i)}
	}

	sort.Sort(a)
	sort.Sort(b)

	var ans = triple{math.MaxInt64, math.MaxInt64, math.MaxInt64}
	var finded = false

	for i := int64(0); i < sizeC; i++ {
		curr, currFinded := TwoSum(a, b, s-c[i].first)
		if !currFinded {
			continue
		} else if !finded {
			ans = curr
			ans.third = c[i].second
			finded = true
		} else {
			curr.third = c[i].second
			//fmt.Println("GOAL:", s, "NUMBERS:", a[curr.first].first, b[curr.second].first, c[curr.third].first)
			ans = Min(ans, curr)
		}
	}

	if !finded {
		fmt.Println(-1)
	} else {
		fmt.Println(ans.first, ans.second, ans.third)
	}
}
