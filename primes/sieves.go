package primes

import "math"

//Eratosthenes resturns an array of int64 that contains the prime numbers up to the max argument
func Eratosthenes(max int64) (primes []int64) {
	if max < 3 {
		return []int64{}
	}

	isPrime := make([]bool, max)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false

	//Cross out multiples of 2
	for i := 4; i < len(isPrime); i += 2 {
		isPrime[i] = false
	}

	limit := int64(math.Sqrt(float64(max)))
	for p := int64(3); p < limit; p += 2 {
		if isPrime[p] {
			for multiple := p * p; multiple < max; multiple = multiple + (2 * p) {
				isPrime[multiple] = false
			}
		}
	}

	for i := range isPrime {
		if isPrime[i] {
			primes = append(primes, int64(i))
		}
	}

	return
}
