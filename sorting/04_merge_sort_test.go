package sorting

import "testing"

func Test_MergeSort(t *testing.T) {
	type want struct {
		first int
		last  int
	}
	tests := []struct {
		name  string
		array []int
		want  want
	}{
		{"First and last swapped", []int{6, 2, 3, 4, 5, 1}, want{1, 6}},
		{"All scrambled", []int{4, 2, 1, 3, 6, 5}, want{1, 6}},
		{"Only two", []int{2, 1}, want{1, 2}},
		{"Only one", []int{1}, want{1, 1}},
	}
	for _, test := range tests {
		lastItem := len(test.array) - 1
		testArray := MergeSort(test.array)
		if testArray[0] != test.want.first || testArray[lastItem] != test.want.last {
			t.Errorf("%q, MergeSort() e:1 is %d and e:n is %d. Order is wrong.", test.name, testArray[0], testArray[lastItem])
		}
	}

}
