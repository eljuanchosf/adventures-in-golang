package numbers

import "testing"

func TestGcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		wantGcd int
	}{
		{"GCD with positive integers", args{24, 30}, 6},
		{"GCD with negative integers", args{-1420, -1572}, -4},
	}
	for _, tt := range tests {
		if gotGcd := Gcd(tt.args.a, tt.args.b); gotGcd != tt.wantGcd {
			t.Errorf("%q. Gcd() = %v, want %v", tt.name, gotGcd, tt.wantGcd)
		}
	}
}
