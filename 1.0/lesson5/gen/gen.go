package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		return
	}

	fileName := args[1]
	fileName += ".go"
	fileContenet := `package main
	
import (
	"bufio"
	"strings"
	"strconv"
)
	
func read_int(sc *bufio.Scanner) int {
	if sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		return num
	}

	return 0
}
	
func read_two_ints(sc *bufio.Scanner) (int, int) {
	if sc.Scan() {
		arr := strings.Fields(sc.Text())
		num1, _ := strconv.Atoi(arr[0])
		num2, _ := strconv.Atoi(arr[1])
		return num1, num2

	}

	return 0, 0
}`

	file, err := os.Create(fileName)

	if err != nil {
		return
	}

	defer file.Close()

	tmpl, err := template.New("main").Parse(fileContenet)
	if err != nil {
		return
	}

	err = tmpl.Execute(file, nil)
	if err != nil {
		return
	}

	fmt.Println("Файл", fileName, "создан")
}
