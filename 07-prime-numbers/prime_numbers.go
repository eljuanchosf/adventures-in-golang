package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ready to begin generating prime numbers. Press CTRL-C to stop.")
	time.Sleep(5 * time.Second)
	var i int64
	// We know that 1 and 2 are prime numbers by default
	i = 3
	for {
		if isPrime(i) {
			fmt.Printf("%d, ", i)
		}
		// Prime numbers can only be odd, since pair numbers are divisible by 2
		i += 2
	}
}

func isPrime(number int64) bool {
	var index int64
	for index = 2; index <= (number / 2); index++ {
		if number%index == 0 {
			return false
		}
	}
	return true
}
