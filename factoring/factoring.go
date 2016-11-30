package factoring

import "math"

//FactorsOf returns an array of int64 with the factors of the number argument
func FactorsOf(number int64) (factors []int64) {

	factors = []int64{}

	if number == 0 {
		return
	}

	for number%2 == 0 {
		factors = append(factors, 2)
		number = number / 2
	}

	factor := float64(3)
	stop := math.Sqrt(float64(number))

	for factor <= stop {
		for number%int64(factor) == 0 {
			factors = append(factors, int64(factor))
			number = number / int64(factor)
			stop = math.Sqrt(float64(number))
		}
		factor += 2
	}
	if number > 1 {
		factors = append(factors, number)
	}
	return
}
