// Modify the previous program such that only multiples of three or five are considered in the sum, e.g. 3, 5, 6, 9, 10, 12, 15 for n=17
package main

import "fmt"

func main() {
	var num int
	fmt.Print("Please enter a number that is multiple of 3 or 5: ")
	fmt.Scanf("%d", &num)
	if isMultipleOf(num, 3) || isMultipleOf(num, 5) {
		result := num + 1
		fmt.Printf("Result of %d + 1 = %d\n", num, result)
	} else {
		fmt.Println("Nooope.")
	}
}

func isMultipleOf(number int, dividend int) bool {
	if (number % dividend) == 0 {
		return true
	}
	return false
}
