package main

import "testing"

func Test_IsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{"2016 is leap", 2016, true},
		{"2020 is leap", 2020, true},
		{"2024 is leap", 2024, true},
		{"2028 is leap", 2028, true},
		{"2000 is leap", 2000, true},
		{"1999 is leap", 1999, false},
		{"1900 is leap", 1900, false},
	}

	for _, tt := range tests {
		if got := isLeapYear(tt.year); got != tt.want {
			t.Errorf("isLeapYear(), year %d, got %v, expected %v", tt.year, got, tt.want)
		}
	}
}
