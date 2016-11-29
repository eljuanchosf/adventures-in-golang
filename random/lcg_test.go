package random

import "testing"
import "github.com/eljuanchosf/adventures-in-golang/sorting"

func Test_GetLCGRandomNumber(t *testing.T) {
	type args struct {
		seed int
		a    int
		b    int
		m    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Generate with valid conditions", args{0, 5, 3, 7}, []int{0, 3, 4, 2, 6, 5}},
		{"Generate breaking the conditions", args{0, 4, 3, 8}, []int{0, 3, 7}},
	}

	for _, test := range tests {
		got := GetLCGRandomNumber(test.args.seed, test.args.a, test.args.b, test.args.m)
		if !sorting.CompareArrays(got, test.want) {
			t.Errorf("%q, GetLCGRandomNumber(), expected %d, got %d", test.name, test.want, got)
		}
	}
}
