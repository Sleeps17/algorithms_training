package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read_int(sc *bufio.Scanner) int64 {
	if sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		return int64(num)
	}

	return 0
}

func read_two_ints(sc *bufio.Scanner) (int64, int64) {
	if sc.Scan() {
		arr := strings.Fields(sc.Text())
		num1, _ := strconv.Atoi(arr[0])
		num2, _ := strconv.Atoi(arr[1])
		return int64(num1), int64(num2)

	}

	return 0, 0
}

func read_slice_ints(sc *bufio.Scanner) []int64 {

	arr := make([]int64, 0)

	if sc.Scan() {
		arrStr := strings.Fields(sc.Text())

		for _, elem := range arrStr {
			num, _ := strconv.Atoi(elem)

			arr = append(arr, int64(num))
		}

		return arr
	}

	return arr
}

func write_int(wr *bufio.Writer, num int64, end string) {
	wr.Write([]byte(strconv.FormatInt(num, 10)))
	wr.Write([]byte(end))
}

func write_two_ints(wr *bufio.Writer, num1, num2 int64, sep, end string) {
	wr.Write([]byte(strconv.FormatInt(num1, 10)))
	wr.Write([]byte(sep))
	wr.Write([]byte(strconv.FormatInt(num2, 10)))
	wr.Write([]byte(end))
}

func write_slice_ints(wr *bufio.Writer, arr []int64, sep, end string) {
	for _, elem := range arr {
		wr.Write([]byte(strconv.FormatInt(elem, 10)))
		wr.Write([]byte(sep))
	}
	wr.Write([]byte(end))
}

func read_string(sc *bufio.Scanner) string {
	if sc.Scan() {
		return sc.Text()
	}

	return ""
}

type point struct {
	x int64
	y int64
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}

	return a
}

func mainJ() {
	in, _ := os.Open("input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 10000000), 10000000)

	n := read_int(sc)
	arr := make([]point, n)
	for i := int64(0); i < n; i++ {
		arr[i].x, arr[i].y = read_two_ints(sc)
	}

	ans := int64(0)
	for i := int64(0); i < n; i++ {
		usedvectors := make(map[point]bool)
		nieb := make([]int64, 0)
		for j := int64(0); j < n; j++ {
			vecx := arr[i].x - arr[j].x
			vecy := arr[i].y - arr[j].y
			veclen := vecx*vecx + vecy*vecy
			nieb = append(nieb, veclen)

			if usedvectors[point{vecx, vecy}] == true {
				ans--
			}
			usedvectors[point{-vecx, -vecy}] = true
		}
		sort.Slice(nieb, func(i, j int) bool {
			return nieb[i] < nieb[j]
		})

		r := int64(0)
		for l := int64(0); l < int64(len(nieb)); l++ {
			for r < int64(len(nieb)) && nieb[l] == nieb[r] {
				r++
			}
			ans += r - l - 1
		}
	}

	out, _ := os.Create("output.txt")
	defer out.Close()
	wr := bufio.NewWriter(out)

	write_int(wr, ans, "\n")
	wr.Flush()
}
