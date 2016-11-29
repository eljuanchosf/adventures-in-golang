package sorting

import "testing"

func Test_InsertionSort(t *testing.T) {
	tests := testCases()
	for _, test := range tests {
		testArray := InsertionSort(test.array)
		if CompareArrays(testArray, test.want) == false {
			t.Errorf("%q, InsertionSort(), expected %v, got %v", test.name, test.want, testArray)
		}

	}
}
