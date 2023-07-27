package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainB() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	dict := make(map[string]int)

	sc := bufio.NewScanner(input)

	for sc.Scan() {
		line := sc.Text()
		arr := strings.Fields(line)

		for _, elem := range arr {
			fmt.Printf("%d ", dict[elem])
			dict[elem]++
		}
	}
	fmt.Println()
}
