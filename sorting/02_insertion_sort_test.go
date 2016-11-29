package sorting

import (
	"testing"

	"github.com/eljuanchosf/adventures-in-golang/shared"
)

func Test_InsertionSort(t *testing.T) {
	tests := shared.TestCases()
	for _, test := range tests {
		testArray := InsertionSort(test.Array)
		if shared.CompareArrays(testArray, test.Want) == false {
			t.Errorf("%q, InsertionSort(), expected %v, got %v", test.Name, test.Want, testArray)
		}

	}
}
