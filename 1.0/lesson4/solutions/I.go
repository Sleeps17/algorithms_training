package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func read_int(sc *bufio.Scanner) int {
	if sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		return num
	}

	return 0
}

func read_two_ints(sc *bufio.Scanner) (int, int) {
	if sc.Scan() {
		arr := strings.Fields(sc.Text())
		num1, _ := strconv.Atoi(arr[0])
		num2, _ := strconv.Atoi(arr[1])
		return num1, num2

	}

	return 0, 0
}

func CheckLegitOfWord(word string) bool {
	countUpperCase := 0

	for _, symb := range word {
		if unicode.IsUpper(symb) {
			countUpperCase++
		}
	}

	if countUpperCase == 0 {
		return false
	} else if countUpperCase == 1 {
		return true
	} else {
		return false
	}
}

func mainI() {
	sc := bufio.NewScanner(os.Stdin)
	n := read_int(sc)
	dict := make(map[string]map[string]bool)

	for i := 0; i < n; i++ {
		sc.Scan()
		word := sc.Text()
		wordLower := strings.ToLower(word)
		if dict[wordLower] == nil {
			dict[wordLower] = make(map[string]bool)
		}
		dict[wordLower][word] = true
	}

	sc.Scan()
	text := sc.Text()

	ans := 0

	for _, word := range strings.Fields(text) {
		wordLower := strings.ToLower(word)
		if _, have := dict[wordLower]; have == true {
			if _, is := dict[wordLower][word]; is == false {
				ans++
			}
		} else {
			stresses := 0
			for _, c := range word {
				if unicode.IsUpper(c) {
					stresses++
				}
			}

			if stresses != 1 {
				ans++
			}
		}

	}

	fmt.Println(ans)
}
