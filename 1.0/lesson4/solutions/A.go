package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainA() {
	rd := bufio.NewReader(os.Stdin)
	var n int
	input, _ := rd.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	n, _ = strconv.Atoi(input)

	dict := make(map[string]string)

	for i := 0; i < n; i++ {
		input, _ = rd.ReadString('\n')
		arr := strings.Fields(input)
		dict[arr[0]] = arr[1]
		dict[arr[1]] = arr[0]
	}

	var word string
	word, _ = rd.ReadString('\n')
	word = strings.TrimSuffix(word, "\n")

	fmt.Println(dict[word])
}
