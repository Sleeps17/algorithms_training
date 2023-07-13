package main

import (
	"fmt"
	"strings"
)

func BringToNormalNumber(s string) string {
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "+7", "8")

	if len(s) != 11 {
		s = "8495" + s
	}

	return s
}
func mainC() {
	s := make([]string, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&s[i])
	}

	for i := 0; i < 4; i++ {
		s[i] = BringToNormalNumber(s[i])
	}

	for i := 1; i < 4; i++ {
		if s[i] == s[0] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
