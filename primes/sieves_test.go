package primes

import (
	"reflect"
	"testing"
)

func Test_Eratosthenes(t *testing.T) {
	type args struct {
		max int64
	}
	tests := []struct {
		name       string
		args       args
		wantPrimes []int64
	}{
		{"Primes up to of 100", args{100}, []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
		{"Primes up to 201", args{201}, []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}},
		{"Primes up to 0", args{0}, []int64{}},
	}
	for _, tt := range tests {
		if gotPrimes := Eratosthenes(tt.args.max); !reflect.DeepEqual(gotPrimes, tt.wantPrimes) {
			t.Errorf("%q. Eratosthenes() = %v, want %v", tt.name, gotPrimes, tt.wantPrimes)
		}
	}
}
