package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainA() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	txt := sc.Text()
	slice := strings.Split(txt, " ")
	flag := true
	for i := 1; i < len(slice); i++ {
		first, _ := strconv.Atoi(slice[i-1])
		second, _ := strconv.Atoi(slice[i])
		if first >= second {
			flag = false
			break
		}
	}

	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
