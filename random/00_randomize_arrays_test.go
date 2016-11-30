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
		got := shared.CopyArray(test.array)
		RandomizeArray1(got)
		if len(got) != len(test.array) {
			t.Errorf("%q, RandomizeArray1(), len of randomized array is %d, expected %d", test.name, len(got), len(test.array))
		}
		if shared.CompareArrays(got, test.array) {
			t.Errorf("%q, RandomizeArray1(), didn't expect %d, got %d", test.name, test.array, got)
		}
	}
}
