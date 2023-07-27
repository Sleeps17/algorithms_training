package main

import (
	"bufio"
	"fmt"
	"os"
)

func MakeDict(dict map[byte]int, word string) {
	for _, symb := range []byte(word) {
		dict[symb]++
	}
}

func MatchDicts(dict1, dict2 map[byte]int) int {
	matches := 0
	for key, value := range dict1 {
		if value == dict2[key] {
			matches++
		}
	}

	return matches
}

func ModifyDict(wDict, sDict map[byte]int, symbol byte, countModifier int) int {
	ans := 0

	if _, have := sDict[symbol]; have != true {
		sDict[symbol] = 0
	}

	if _, have := wDict[symbol]; have == true && sDict[symbol] == wDict[symbol] {
		ans = -1
	}

	sDict[symbol] += countModifier

	if _, have := wDict[symbol]; have == true && wDict[symbol] == sDict[symbol] {
		ans = 1
	}

	return ans
}

func mainH() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	lenW, lenS := read_two_ints(sc)
	sc.Scan()
	w := sc.Text()
	sc.Scan()
	s := sc.Text()

	wDict := make(map[byte]int)
	MakeDict(wDict, w)
	newS := []byte(s)
	sDict := make(map[byte]int)
	MakeDict(sDict, string(newS[:lenW]))
	matchingLettres := MatchDicts(wDict, sDict)
	ans := 0
	if matchingLettres == len(wDict) {
		ans++
	}

	for i := lenW; i < lenS; i++ {
		matchingLettres += ModifyDict(wDict, sDict, newS[i-lenW], -1)
		matchingLettres += ModifyDict(wDict, sDict, newS[i], 1)

		if matchingLettres == len(wDict) {
			ans++
		}
	}

	fmt.Println(ans)
}
