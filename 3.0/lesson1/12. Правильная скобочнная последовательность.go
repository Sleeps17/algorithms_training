package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	var seq string
	if sc.Scan() {
		seq = sc.Text()
	}

	if PSP([]rune(seq)) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
