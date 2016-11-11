package main

import "testing"

func TestIsPrimeNumber(t *testing.T) {
	type args struct {
		number int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1 is prime", args{1}, true},
		{"3 is prime", args{3}, true},
		{"10 is not prime", args{10}, false},
		{"7 is prime", args{7}, true},
		{"198 is not prime", args{198}, false},
		{"199 is prime", args{199}, true},
		{"263201 is prime", args{263201}, true},
	}
	for _, tt := range tests {
		if got := isPrime(tt.args.number); got != tt.want {
			t.Errorf("%q. isPrime() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
