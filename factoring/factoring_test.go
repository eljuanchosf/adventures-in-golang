package factoring

import (
	"reflect"
	"testing"
)

func Test_factorsOf(t *testing.T) {
	type args struct {
		number int64
	}
	tests := []struct {
		name        string
		args        args
		wantFactors []int64
	}{
		{"Factors of 0", args{0}, []int64{}},
		{"Factors of 204", args{204}, []int64{2, 2, 3, 17}},
		{"Factors of a... large number.", args{123456789012345678}, []int64{2, 3, 3, 3, 21491747, 106377431}},
	}
	for _, tt := range tests {
		if gotFactors := FactorsOf(tt.args.number); !reflect.DeepEqual(gotFactors, tt.wantFactors) {
			t.Errorf("%q. factorsOf() = %v, want %v", tt.name, gotFactors, tt.wantFactors)
		}
	}
}
