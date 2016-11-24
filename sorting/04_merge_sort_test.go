package sorting

import "testing"

func Test_MergeSort(t *testing.T) {
	tests := testCases()
	for _, test := range tests {
		testArray := MergeSort(test.array)
		if compareArrays(testArray, test.want) == false {
			t.Errorf("%q, MergeSort(), expected %v, got %v", test.name, test.want, testArray)
		}
	}
}
