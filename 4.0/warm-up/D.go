package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainD() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	var str1, str2 string

	sc.Scan()
	str1 = sc.Text()
	sc.Scan()
	str2 = sc.Text()

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, char := range str1 {
		m1[char]++
	}

	for _, char := range str2 {
		m2[char]++
	}

	if len(m1) != len(m2) {
		fmt.Println("NO")
		return
	}

	for key, value := range m1 {

		if m2[key] != value {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
