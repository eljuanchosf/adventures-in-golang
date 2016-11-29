package sorting

import "testing"

func Test_HeapSort(t *testing.T) {
	tests := testCases()
	for _, test := range tests {
		testArray := HeapSort(test.array)
		if CompareArrays(testArray, test.want) == false {
			t.Errorf("%q, HeapSort(), expected %v, got %v", test.name, test.want, testArray)
		}

	}
}
