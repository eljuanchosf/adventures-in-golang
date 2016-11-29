package random

//GetLCGRandomNumber generates an array of pseudo random numbers.
// Linear Congruential Generators is a way to generate pseudo random numbers based on a seed. It uses the formula:
//  X(n-1) = X(n) * A + B (mod M)
func GetLCGRandomNumber(seed, a, b, m int) (array []int) {
	sequenceSeed := seed
	array = append(array, seed)
	result := (seed*a + b) % m
	for seed != result && sequenceSeed != result {
		array = append(array, result)
		sequenceSeed = result
		result = (sequenceSeed*a + b) % m
	}
	return
}
