package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ParseString(str string) (string, int64) {

	arr := strings.Fields(str)

	cnt, _ := strconv.Atoi(arr[len(arr)-1])

	return strings.Join(arr[:len(arr)-1], " "), int64(cnt)
}

func mainD() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	votes := make(map[string]int64)
	partys := make([]string, 0, 20)
	sum := int64(0)

	for sc.Scan() {
		party, cnt := ParseString(sc.Text())

		votes[party] += cnt
		partys = append(partys, party)

		sum += cnt
	}

	places := make(map[string]int64)
	doubles := make(map[float64]string)
	curr_sum := int64(0)
	for _, party := range partys {
		places[party] = int64((float64(votes[party]) / float64(sum)) * float64(450))
		curr_sum += places[party]
		doubles[float64(votes[party])/(float64(sum)/float64(450))-float64(places[party])] = party
	}

	if curr_sum < 450 {
		var keys []float64
		for key := range doubles {
			keys = append(keys, key)
		}

		slices.Sort(keys)

		for curr_sum < 450 {
			for i := len(keys) - 1; i >= 0 && curr_sum < 450; i-- {
				places[doubles[keys[i]]]++
				curr_sum++
			}
		}
	}

	for _, party := range partys {
		fmt.Println(party, places[party])
	}

}
