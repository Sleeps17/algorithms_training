package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main3() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	var expr string
	if sc.Scan() {
		expr = sc.Text()
	}

	expr = strings.ReplaceAll(expr, " ", "")
	ans := CalculateExper(expr)
	fmt.Println(ans)
}
