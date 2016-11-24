package sorting

import "testing"

func Test_BubleSort(t *testing.T) {
	tests := testCases()
	for _, test := range tests {
		testArray := BubleSort(test.array)
		if compareArrays(testArray, test.want) == false {
			t.Errorf("%q, BubleSort(), expected %v, got %v", test.name, test.want, testArray)
		}
	}
}
