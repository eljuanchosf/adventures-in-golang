package numbers

//Lcm calculates the Least common multiple between numbers a and between
func Lcm(a, b int) int {
	return (a / Gcd(a, b)) * b
}
