package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainD() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000000), 10000000)

	var str string
	if sc.Scan() {
		str = sc.Text()
	}

	cnt := int64(0)
	i := int64(0)

	for i < int64(len(str)) && cnt >= int64(0) {
		if str[i] == '(' {
			cnt++
		} else {
			cnt--
		}
		i++
	}

	if cnt == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
