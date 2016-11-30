package numbers

import "testing"

func TestLcm(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"LCM for postivie integers", args{12, 15}, 60},
	}
	for _, tt := range tests {
		if got := Lcm(tt.args.a, tt.args.b); got != tt.want {
			t.Errorf("%q. Lcm() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
