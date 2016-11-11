package main

import "testing"

func Test_isMultipleOf(t *testing.T) {
	type args struct {
		number   int
		dividend int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Not a multiple",
			args{11, 5},
			false,
		},
		{"A multiple",
			args{10, 5},
			true,
		},
	}
	for _, tt := range tests {
		if got := isMultipleOf(tt.args.number, tt.args.dividend); got != tt.want {
			t.Errorf("%q. isMultipleOf() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
