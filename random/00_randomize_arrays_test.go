package random

import (
	"testing"

	"github.com/eljuanchosf/adventures-in-golang/shared"
)

func Test_RandomizeArray1(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		dontWant []int
	}{
		{"Array 1", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
	}

	for _, test := range tests {
		originalArray := shared.CopyArray(test.array)
		RandomizeArray1(test.array)
		if len(originalArray) != len(test.array) {
			t.Errorf("%q, RandomizeArray1(), len of randomized array is %d, expected %d", test.name, len(test.array), len(originalArray))
		}
		if shared.CompareArrays(originalArray, test.array) {
			t.Errorf("%q, RandomizeArray1(), didn't expect %d, got %d", test.name, originalArray, test.array)
		}
	}
}
