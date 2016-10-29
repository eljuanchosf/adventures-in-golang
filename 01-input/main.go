// Write a program that asks the user for her name and greets her with her name.
package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Please enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello, %s!\n", name)
}
