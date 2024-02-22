package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInt(sc *bufio.Scanner) int {
	var nums []string
	if sc.Scan() {
		nums = strings.Fields(sc.Text())
	}

	num, _ := strconv.Atoi(nums[0])

	return num
}

func WriteString(w *bufio.Writer, str string, end string) {
	w.WriteString(str)
	w.WriteString(end)
}

func WriteArrString(w *bufio.Writer, arr []string, sep, end string) {
	for i, elem := range arr {
		w.WriteString(elem)

		if i == len(arr)-1 {
			w.WriteString(end)
		} else {
			w.WriteString(sep)
		}
	}

	if len(arr) == 0 {
		w.WriteString("\n")
	}
}

func mainE() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 1000000)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := ReadInt(sc)

	arr := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		arr[i] = sc.Text()
	}

	WriteString(w, "Initial array:", "\n")
	WriteArrString(w, arr, ", ", "\n")

	maxLen := len(arr[0])
	for i := 1; i < n; i++ {
		if len(arr[i]) > maxLen {
			maxLen = len(arr[i])
		}
	}

	buckets := make([][]string, 10)

	for i := 0; i < maxLen; i++ {
		WriteString(w, "**********", "\n")
		WriteString(w, fmt.Sprintf("Phase %d\n", i+1), "")

		for j := 0; j < 10; j++ {
			buckets[j] = make([]string, 0)
		}

		for j := 0; j < n; j++ {
			num, _ := strconv.Atoi(string(arr[j][len(arr[j])-i-1]))
			buckets[num] = append(buckets[num], arr[j])
		}

		for j := 0; j < 10; j++ {
			WriteString(w, fmt.Sprintf("Bucket %d: ", j), "")
			if len(buckets[j]) == 0 {
				WriteString(w, "empty", "\n")
			} else {
				WriteArrString(w, buckets[j], ", ", "\n")
			}
		}

		index := 0
		for j := 0; j < 10; j++ {
			for k := 0; k < len(buckets[j]); k++ {
				arr[index] = buckets[j][k]
				index++
			}
		}
	}

	WriteString(w, "**********", "\n")
	WriteString(w, "Sorted array:", "\n")
	WriteArrString(w, arr, ", ", "\n")
}
