package main

import "testing"

func TestGenerateMultiplicationTable(t *testing.T) {
	type args struct {
		multiplier int
	}
	tests := []struct {
		name     string
		args     args
		position int
		want     int
	}{
		{"Table of 4", args{4}, 5, 20},
	}
	for _, tt := range tests {
		if got := generateMultiplicationTable(tt.args.multiplier)[tt.position]; got != tt.want {
			t.Errorf("%q. GenerateMuliplicationTable(), position %d = %v, want %v", tt.name, tt.position, got, tt.want)
		}
	}
}
