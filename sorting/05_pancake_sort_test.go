package sorting

import (
	"testing"

	"github.com/eljuanchosf/adventures-in-golang/shared"
)

func Test_PancakeSort(t *testing.T) {
	tests := shared.TestCases()
	for _, test := range tests {
		sortedArray := PancakeSort(test.Array)
		if shared.CompareArrays(sortedArray, test.Want) == false {
			t.Errorf("%q, PancakeSort(), expected %v, got %v.", test.Name, test.Want, sortedArray)
		}
	}

}
