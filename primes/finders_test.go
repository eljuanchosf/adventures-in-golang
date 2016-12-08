package primes

import "testing"

func Test_Fermat(t *testing.T) {
	tests := []struct {
		number int64
		want   bool
	}{
		{2, true},
		{100, false},
		//		{3354895165247, false},
		{11, true},
	}

	for _, test := range tests {
		if got := Fermat(test.number); got != test.want {
			t.Errorf("Fermat(%d), got %v, expected %v", test.number, got, test.want)
		}
	}
}
