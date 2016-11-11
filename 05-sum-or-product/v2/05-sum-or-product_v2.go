// Package v2 Write a program that asks the user for a number n and gives him the possibility to choose between computing the sum and computing the product of 1,…,n.
package v2

import "fmt"

func main() {
	num1 := askForNumber("Please enter a number: ")
	opt := askForString("Do you want to compute the (P)roduct or the (S)um? ")
	if opt == "P" {
		result := multiplyRange(1, num1)
		fmt.Printf("The result of the operation you chose is: %d\n", result)
	} else if opt == "S" {
		result := sumRange(1, num1)
		fmt.Printf("The result of the operation you chose is: %d\n", result)
	} else {
		fmt.Println("Not valid")
	}
}

func askForNumber(message string) int {
	var num int
	fmt.Printf(message)
	fmt.Scanf("%d", &num)
	return num
}

func askForString(message string) string {
	var str string
	fmt.Printf(message)
	fmt.Scanf("%s", &str)
	return str
}

func sumRange(start int, end int) int {
	ac := end * (start + end) / 2
	return ac
}

func multiplyRange(start int, end int) int {
	ac := 1
	for i := start; i <= end; i++ {
		ac = ac * i
	}
	return ac
}
