package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadInt(sc *bufio.Scanner) int64 {
	num := 0
	if sc.Scan() {
		num, _ = strconv.Atoi(sc.Text())
	}

	return int64(num)
}

func ReadTwoInts(sc *bufio.Scanner) (int64, int64) {
	num1, num2 := 0, 0

	if sc.Scan() {
		seq := strings.Fields(sc.Text())

		num1, _ = strconv.Atoi(seq[0])
		num2, _ = strconv.Atoi(seq[1])
	}

	return int64(num1), int64(num2)
}

func ReadSliceInts(sc *bufio.Scanner) []int64 {

	if sc.Scan() {
		seq := strings.Fields(sc.Text())

		arr := make([]int64, 0, len(seq))

		for i := 0; i < len(seq); i++ {
			num, _ := strconv.Atoi(seq[i])
			arr = append(arr, int64(num))
		}

		return arr
	}

	return nil
}

func ReadString(sc *bufio.Scanner) string {

	if sc.Scan() {
		return sc.Text()
	}

	return ""
}

func WriteInt(wr *bufio.Writer, num int64, end string) {
	wr.Write([]byte(strconv.FormatInt(num, 10)))
	wr.Write([]byte(end))
}

func WriteTwoInts(wr *bufio.Writer, num1, num2 int64, sep, end string) {
	WriteInt(wr, num1, sep)
	WriteInt(wr, num2, end)
}

func WriteSliceInts(wr *bufio.Writer, arr []int64, sep, end string) {
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			WriteInt(wr, arr[i], end)
		}
		WriteInt(wr, arr[i], sep)
	}

}

func WriteString(wr *bufio.Writer, str, end string) {
	wr.Write([]byte(str))
	wr.Write([]byte(end))
}

func lower_bound(arr []int64, findelem int64) int64 {
	l := 0
	r := len(arr)

	for l < r {
		m := (l + r) / 2

		if arr[m] < findelem {
			l = m + 1
		} else {
			r = m
		}
	}

	return int64(l)
}

func upper_bound(arr []int64, findelem int64) int64 {
	l := 0
	r := len(arr)

	for l < r {
		m := (l + r) / 2

		if arr[m] <= findelem {
			l = m + 1
		} else {
			r = m
		}
	}

	return int64(r)
}

func mainI() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 1000000), 1000000)
	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	data := ReadSliceInts(sc)
	n, r, c := data[0], data[1], data[2]
	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		arr[i] = ReadInt(sc)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	left := int64(0)
	right := int64(math.MaxInt64)
	for left < right {
		mid := (left + right) / 2

		cnt := int64(0)
		for i := c - 1; i < n; {
			if arr[i]-arr[i-c+1] <= mid {
				cnt++
				i += c
			} else {
				i++
			}
		}

		if cnt < r {
			left = mid + 1
		} else {
			right = mid
		}
	}

	WriteInt(wr, left, "\n")
	wr.Flush()
}
