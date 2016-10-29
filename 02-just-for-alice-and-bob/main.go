// Modify the previous program such that only the users Alice and Bob are greeted with their names.
package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Please enter your name: ")
	fmt.Scanf("%s", &name)
	if name == "Alice" || name == "Bob" {
		fmt.Printf("Hello, %s!\n", name)
	}
}
