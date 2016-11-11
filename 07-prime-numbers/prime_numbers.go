package main

import "fmt"

func main() {
	for i := 0; i < 32638; i++ {
		if isPrime(i) {
			fmt.Printf("%d\n", i)
		}
	}
}

func isPrime(number int) bool {
	for index := 2; index <= (number / 2); index++ {
		if number%index == 0 {
			return false
		}
	}
	return true
}
