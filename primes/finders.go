package primes

import (
	"math"
	"math/rand"
	"time"
)

//Fermat returns true if the argument number is a prime number
func Fermat(number int64) (isPrime bool) {
	if number > 0 && number < 3 {
		return true
	}
	if number%2 == 0 {
		return false
	}

	for iteration := 1; iteration < 100; iteration++ {
		r := random(2, number)
		power := number - 1
		n := int64(math.Pow(float64(r), float64(power))) % number
		if n != 1 {
			return false
		}
	}
	return true
}

func random(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
