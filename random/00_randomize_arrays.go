package random

import (
	"math/rand"

	"github.com/eljuanchosf/adventures-in-golang/shared"
)

//RandomizeArray1 randomizes an array and returns the result
func RandomizeArray1(array []int) []int {
	elementCount := len(array) - 1
	for i := 0; i < elementCount; i++ {
		j := rand.Intn(elementCount)
		shared.SwapInt(array, i, j)
	}
	return array
}
