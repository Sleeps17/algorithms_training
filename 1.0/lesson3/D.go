package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func mainD() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	arr := strings.Fields(string(data))
	set := make(map[string]bool)

	for _, elem := range arr {
		set[elem] = true
	}

	ioutil.WriteFile("output.txt", []byte(strconv.Itoa(len(set))), 0644)
}
