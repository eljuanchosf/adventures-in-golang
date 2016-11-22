package sorting

import "testing"

func Test_swapInt(t *testing.T) {
	array := []int{2, 1}
	swapInt(array, 0, 1)
	if array[0] != 1 {
		t.Error("swapInt is not working correctly")
	}
}

func Test_findSmallestElement(t *testing.T) {
	type args struct {
		array      []int
		startIndex int
	}
	tests := []struct {
		name           string
		args           args
		wantedPosition int
	}{
		{"First and last swapped", args{[]int{6, 2, 3, 4, 5, 1}, 0}, 5},
		{"All scrambled", args{[]int{4, 2, 0, 3, 6, 5}, 3}, 3},
	}
	for _, test := range tests {
		position := findSmallestElement(test.args.array, test.args.startIndex)
		if position != test.wantedPosition {
			t.Errorf("%q, findSmallesElement(), expected (p:%d), got (p:%d).", test.name, test.wantedPosition, position)
		}
	}
}
