package sorting

import "testing"

func Test_SelectiveSort(t *testing.T) {
	tests := testCases()
	for _, test := range tests {
		testArray := SelectiveSort(test.array)
		if compareArrays(testArray, test.want) == false {
			t.Errorf("%q, SelectiveSort(), expected %v, got %v", test.name, test.want, testArray)
		}

	}
}
