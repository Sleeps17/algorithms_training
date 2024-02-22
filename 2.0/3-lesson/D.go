package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInt(sc *bufio.Scanner) int {
	sc.Scan()

	num, _ := strconv.Atoi(sc.Text())

	return num
}

func StringToMap(str string) map[int]bool {
	strNums := strings.Fields(str)

	res := make(map[int]bool)

	for _, strNum := range strNums {
		num, _ := strconv.Atoi(strNum)

		res[num] = true
	}

	return res
}

func CrossSets(m1, m2 *map[int]bool) {

	//	fmt.Println("\t\t", m1)
	//	fmt.Println("\t\t", m2)

	for key := range *m1 {
		(*m1)[key] = (*m1)[key] && (*m2)[key]
	}

	//	fmt.Println("\t\t", m1)
	//	fmt.Println("\t\t", m2)

}

func UnionSets(m1, m2 *map[int]bool) {

	//	fmt.Println("\t\t", m1)
	//	fmt.Println("\t\t", m2)

	for key := range *m2 {
		(*m1)[key] = (*m1)[key] || (*m2)[key]
	}

	// fmt.Println("\t\t", m1)
	// fmt.Println("\t\t", m2)
}

func mainD() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	yes := make(map[int]bool)
	no := make(map[int]bool)

	n := ReadInt(sc)
	for i := 1; i <= n; i++ {
		yes[i] = true
	}

	var m map[int]bool

	for sc.Scan() && sc.Text() != "HELP" {
		str := sc.Text()

		if str != "YES" && str != "NO" {
			m = StringToMap(str)
		}

		if str == "YES" {
			//fmt.Println("YES ACTION")
			CrossSets(&yes, &m)
		}

		if str == "NO" {
			//fmt.Println("NO ACTION")
			UnionSets(&no, &m)
		}
	}

	for i := 1; i <= n; i++ {
		if yes[i] && !no[i] {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
