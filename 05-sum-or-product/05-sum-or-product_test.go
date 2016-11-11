package sum_or_product

import "testing"

func Test_sumRange(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Sum of 1 to 10", args{1, 10}, 55},
	}
	for _, tt := range tests {
		if got := sumRange(tt.args.start, tt.args.end); got != tt.want {
			t.Errorf("%q. sumRange() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_multiplyRange(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Product of 1 to 10", args{1, 10}, 3628800},
	}
	for _, tt := range tests {
		if got := multiplyRange(tt.args.start, tt.args.end); got != tt.want {
			t.Errorf("%q. multiplyRange() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestTestSumRange(t *testing.T) {
	type args struct {
		t *testing.T
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		TestSumRange(tt.args.t)
	}
}

func TestTestMultiplyRange(t *testing.T) {
	type args struct {
		t *testing.T
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		TestMultiplyRange(tt.args.t)
	}
}
