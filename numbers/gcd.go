package numbers

//Gcd calculates the greatest common divisor between two numbers
func Gcd(a, b int) int {
	for b != 0 {
		remainder := a % b
		a = b
		b = remainder
	}
	return a
}
