package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadArr(sc *bufio.Scanner) []int {
	arr := make([]int, 0, 10)
	var nums []string
	if sc.Scan() {
		nums = strings.Fields(sc.Text())
	}

	for _, elem := range nums {
		num, _ := strconv.Atoi(elem)
		arr = append(arr, num)
	}

	return arr
}

func WriteArr(w *bufio.Writer, arr []int, sep string, end string) {
	for i, elem := range arr {
		str := strconv.Itoa(elem)

		w.Write([]byte(str))
		if i == len(arr)-1 {
			w.Write([]byte(end))
		} else {
			w.Write([]byte(sep))
		}
	}
}

func mergeC(arr []int, l, m, r int) {
	i := l
	j := m + 1

	res := make([]int, 0, r-l+1)

	for i < m+1 && j <= r {
		if arr[i] < arr[j] {
			res = append(res, arr[i])
			i++
		} else {
			res = append(res, arr[j])
			j++
		}
	}

	for i < m+1 {
		res = append(res, arr[i])
		i++
	}

	for j < r {
		res = append(res, arr[j])
		j++
	}
	for k := range res {
		arr[l+k] = res[k]
	}
}

func mergeSort(arr []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		mergeC(arr, l, m, r)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 1000000)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := ReadInt(sc)
	arr := ReadArr(sc)

	mergeSort(arr, 0, n-1)
	WriteArr(w, arr, " ", "\n")
}
