package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main1() {
	sc := bufio.NewScanner(os.Stdin)
	var s Stack[int]
	s.Create()

	for sc.Scan() {
		arr := strings.Fields(sc.Text())
		command := arr[0]

		switch command {
		case "push":
			fmt.Println("ok")
			num, _ := strconv.Atoi(arr[1])
			s.Push(num)
		case "pop":
			if s.Empty() {
				fmt.Println("error")
			} else {
				fmt.Println(s.Top())
				s.Pop()
			}
		case "back":
			if s.Empty() {
				fmt.Println("error")
			} else {
				fmt.Println(s.Top())
			}
		case "size":
			fmt.Println(s.Size())
		case "clear":
			s.Clear()
			fmt.Println("ok")
		default:
			fmt.Println("bye")
			return
		}
	}
}
