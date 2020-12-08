package main

import "fmt"

func main() {
	var input1, input2 string
	var constant string = "CONST"
	fmt.Scan(&input1)
	fmt.Scan(&input2)
	if input1 == "" {
		println("An Empty String")
	}

	if input1 != "" {
		println("Not an Empty String")
	}

	if constant == "CONST" {
		println("1")
	}
	if constant != "CONST" {
		println("2")
	}
	if constant == constant {
		println("3")
	}
	if constant != constant {
		println("4")
	}
	if input1 == input2 {
		println("5")
	}
	if input1 != input2 {
		println("6")
	}
	if input2 == constant {
		println("7")
	}
	if input2 != constant {
		println("8")
	}
}
