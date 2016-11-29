package sorting

import "testing"

func Test_PancakeSort(t *testing.T) {
	tests := testCases()
	for _, test := range tests {
		sortedArray := PancakeSort(test.array)
		if CompareArrays(sortedArray, test.want) == false {
			t.Errorf("%q, PancakeSort(), expected %v, got %v.", test.name, test.want, sortedArray)
		}
	}

}
