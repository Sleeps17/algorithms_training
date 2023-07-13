package main

import "fmt"

func mainA() {
	var troom, tcond int
	var mode string

	fmt.Scan(&troom, &tcond, &mode)
	if mode == "auto" {
		fmt.Println(tcond)
	} else if troom < tcond && mode == "heat" {
		fmt.Println(tcond)
	} else if troom == tcond && mode == "fan" {
		fmt.Println(tcond)
	} else if troom > tcond && mode == "freeze" {
		fmt.Println(tcond)
	} else {
		fmt.Println(troom)
	}
}
