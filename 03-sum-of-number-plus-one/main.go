// Write a program that asks the user for a number n and prints the sum of the numbers 1 to n
package main

import "fmt"

func main() {
	var num int
	fmt.Print("Please enter a number: ")
	fmt.Scanf("%d", &num)
	result := num + 1
	fmt.Printf("Result of %d + 1 = %d\n", num, result)
}
