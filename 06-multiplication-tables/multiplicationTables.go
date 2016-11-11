// Write a program that prints a multiplication table for numbers up to 12.
package main

import "fmt"

func main() {
	for i := 1; i < 13; i++ {
		fmt.Printf("\nTable of %d\n", i)
		table := generateMultiplicationTable(i)
		for r := 0; r <= len(table)-1; r++ {
			fmt.Printf("%d * %d = %d\n", i, r, table[r])
		}
	}
}

func generateMultiplicationTable(multiplier int) []int {
	var table []int
	for i := 0; i < 11; i++ {
		table = append(table, i*multiplier)
	}
	return table
}
