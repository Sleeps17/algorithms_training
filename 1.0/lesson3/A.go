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
	input, _ := rd.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	arr := strings.Split(input, " ")

	st := NewSet(10)
	for _, elem := range arr {
		num, _ := strconv.Atoi(elem)
		st.add(num)
	}

	fmt.Println(st.setcount)
}
