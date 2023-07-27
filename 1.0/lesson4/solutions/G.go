package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func deposite(bank map[string]int, name string, sum int) {
	bank[name] += sum
}

func withdraw(bank map[string]int, name string, sum int) {
	bank[name] -= sum
}

func balance(bank map[string]int, name string) {
	bal, err := bank[name]

	if err == false {
		fmt.Println("ERROR")
	} else {
		fmt.Println(bal)
	}
}

func transfer(bank map[string]int, name1, name2 string, sum int) {
	bank[name1] -= sum
	bank[name2] += sum
}

func income(bank map[string]int, p int) {
	for name, bal := range bank {
		if bal <= 0 {
			continue
		}

		per := float64(p)
		dbal := float64(bal)
		bal = int(dbal * (1.0 + per/100))
		bank[name] = bal
	}
}

func mainG() {
	bank := make(map[string]int)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input := sc.Text()

		arr := strings.Fields(input)

		if arr[0] == "DEPOSIT" {
			name := arr[1]
			sum, _ := strconv.Atoi(arr[2])
			deposite(bank, name, sum)
		} else if arr[0] == "WITHDRAW" {
			name := arr[1]
			sum, _ := strconv.Atoi(arr[2])
			withdraw(bank, name, sum)
		} else if arr[0] == "BALANCE" {
			name := arr[1]
			balance(bank, name)
		} else if arr[0] == "TRANSFER" {
			name1, name2 := arr[1], arr[2]
			sum, _ := strconv.Atoi(arr[3])
			transfer(bank, name1, name2, sum)
		} else if arr[0] == "INCOME" {
			p, _ := strconv.Atoi(arr[1])
			income(bank, p)
		}
	}
}
