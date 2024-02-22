package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadInt() int64 {
	num := 0
	if sc.Scan() {
		num, _ = strconv.Atoi(sc.Text())
	}

	return int64(num)
}

func ReadTwoInts() (int64, int64) {
	num1, num2 := 0, 0

	if sc.Scan() {
		seq := strings.Fields(sc.Text())

		num1, _ = strconv.Atoi(seq[0])
		num2, _ = strconv.Atoi(seq[1])
	}

	return int64(num1), int64(num2)
}

func ReadSliceInts() []int64 {

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

func ReadString() string {

	if sc.Scan() {
		return sc.Text()
	}

	return ""
}

func WriteInt(num int64, end string) {
	wr.Write([]byte(strconv.FormatInt(int64(num), 10)))
	wr.Write([]byte(end))
}

func WriteTwoInts(num1, num2 int64, sep, end string) {
	WriteInt(num1, sep)
	WriteInt(num2, end)
}

func WriteSliceInts(arr []int64, sep, end string) {
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			WriteInt(arr[i], end)
			break
		}
		WriteInt(arr[i], sep)
	}

}

func WriteString(str, end string) {
	wr.Write([]byte(str))
	wr.Write([]byte(end))
}

func ReadSliceEvents(d int64) []event {
	events := make([]event, 0)
	if sc.Scan() {
		seq := strings.Fields(sc.Text())

		for i := 0; i < len(seq); i++ {
			num, _ := strconv.Atoi(seq[i])
			events = append(events, event{int64(num), -1})
			events = append(events, event{int64(num) + d, 1})
		}
	}

	return events
}

type event struct {
	Time int64
	Type int64
}

var in, _ = os.Open("{input.txt")
var out, _ = os.Create("output.txt")
var sc = bufio.NewScanner(in)
var wr = bufio.NewWriter(out)

type cash struct {
	Time   int64
	Type   int64
	Number int64
}

func mainE() {
	sc.Buffer(make([]byte, 10000000), 10000000)
	defer in.Close()
	defer out.Close()

	n := ReadInt()
	cnt := int64(0)
	cashes := make([]cash, 0, 2*n)
	for i := int64(0); i < n; i++ {
		time := ReadSliceInts()
		begin := 60*time[0] + time[1]
		end := 60*time[2] + time[3]
		if begin == end {
			cnt++
			continue
		}
		cashes = append(cashes, cash{begin, 1, i + 1})
		cashes = append(cashes, cash{end, 2, i + 1})
	}
	n -= int64(cnt)

	sort.Slice(cashes, func(i, j int) bool {
		if cashes[i].Time == cashes[j].Time {
			return cashes[i].Type > cashes[j].Type
		}

		return cashes[i].Time < cashes[j].Time
	})

	openedCashes := make(map[int64]bool)
	cnt = int64(0)
	st := int64(0)
	end := int64(0)
	ans := int64(0)

	for _, ch := range cashes {
		if ch.Type == 1 {
			cnt++
			openedCashes[ch.Number] = true
			if cnt == n {
				st = ch.Time
			}
		} else {
			if !openedCashes[ch.Number] {
				continue
			}

			if cnt == n {
				end = ch.Time
				ans += end - st
			}

			cnt--
		}
	}

	for _, ch 
}
