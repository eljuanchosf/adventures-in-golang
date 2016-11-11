// Print the 20 next leap years
package main

import (
	"fmt"
	"time"
)

func main() {
	howMany := 20
	year := time.Now().Year()
	i := 1
	for {
		if isLeapYear(year) {
			fmt.Printf("%02d - %d is a leap year\n", i, year)
			if i == howMany {
				break
			}
			i++
		}
		year++
	}
}

func isLeapYear(year int) bool {
	isLeap := false
	if year%4 != 0 {
		isLeap = false
	} else if year%100 != 0 {
		isLeap = true
	} else if year%400 != 0 {
		isLeap = false
	} else {
		isLeap = true
	}
	return isLeap
}
