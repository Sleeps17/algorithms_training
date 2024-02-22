package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainC() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)
	sc.Scan()
	s := sc.Text()
	z := zFunction(s)
	for i := 0; i < len(z); i++ {
		fmt.Printf("%d ", z[i])
	}
	fmt.Println()
}
