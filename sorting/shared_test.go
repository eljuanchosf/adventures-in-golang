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
		endIndex   int
	}
	tests := []struct {
		name           string
		args           args
		wantedPosition int
	}{
		{"Start index is 0, end index is 5", args{[]int{6, 2, 3, 4, 5, 1}, 0, 5}, 5},
		{"Start index is 1, end index is 4", args{[]int{4, 2, 0, 3, 6, 5}, 1, 4}, 2},
	}
	for _, test := range tests {
		position := findSmallestElement(test.args.array, test.args.startIndex, test.args.endIndex)
		if position != test.wantedPosition {
			t.Errorf("%q, findSmallesElement(), expected (p:%d), got (p:%d).", test.name, test.wantedPosition, position)
		}
	}
}

func Test_findBiggestElement(t *testing.T) {
	type args struct {
		array      []int
		startIndex int
		endIndex   int
	}
	tests := []struct {
		name           string
		args           args
		wantedPosition int
	}{
		{"Start index is 0, end index is 5", args{[]int{6, 2, 3, 4, 5, 1}, 0, 5}, 0},
		{"Start index is 1, end index is 4", args{[]int{4, 2, 0, 3, 6, 5}, 1, 4}, 4},
	}
	for _, test := range tests {
		position := findBiggestElement(test.args.array, test.args.startIndex, test.args.endIndex)
		if position != test.wantedPosition {
			t.Errorf("%q, findBiggestElement(), expected (p:%d), got (p:%d).", test.name, test.wantedPosition, position)
		}
	}
}

func Test_CompareArrays(t *testing.T) {
	type args struct {
		arrayA []int
		arrayB []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Arrays are equal", args{[]int{6, 2, 3, 4, 5, 1}, []int{6, 2, 3, 4, 5, 1}}, true},
		{"Arrays are not equal", args{[]int{1, 1, 1, 1, 6, 5}, []int{4, 2, 0, 3, 6, 5}}, false},
		{"Arrays are empty", args{[]int{}, []int{}}, true},
		{"Arrays are null", args{}, true},
		{"ArrayB is null", args{arrayA: []int{}}, false},
	}
	for _, test := range tests {
		if got := CompareArrays(test.args.arrayA, test.args.arrayB); got != test.want {
			t.Errorf("%q, CompareArrays(), expected %v, got %v.", test.name, test.want, got)
		}
	}
}

func Test_flipArray(t *testing.T) {
	type args struct {
		array      []int
		startIndex int
		endIndex   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Flip from 2 to 4", args{[]int{1, 2, 3, 4, 5, 6}, 2, 4}, []int{1, 2, 5, 4, 3, 6}},
		{"Flip from 0 to 5", args{[]int{1, 2, 3, 4, 5, 6}, 0, 5}, []int{6, 5, 4, 3, 2, 1}},
	}
	for _, test := range tests {
		array := make([]int, len(test.args.array))
		copy(array, test.args.array)
		flipArray(array, test.args.startIndex, test.args.endIndex)
		if CompareArrays(array, test.want) == false {
			t.Errorf("%q, flipArray(), original %v, expected %v, got %v", test.name, test.args.array, test.want, array)
		}
	}
}
