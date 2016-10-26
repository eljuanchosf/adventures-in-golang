package addition_test

import (
	"testing"

	"github.com/eljuanchosf/adventures-in-golang/01-testing-101"
)

func TestAddSomeNumbers(t *testing.T) {
	var sumTests = []struct {
		a        int // input a
		b        int // input b
		expected int // expected result
	}{
		{1, 1, 2},
		{10, 20, 30},
		{1938, 1333, 3271},
	}

	for _, test := range sumTests {
		v := addition.AddSomeNumbers(test.a, test.b)
		if v != test.expected {
			t.Errorf("Add some numbers is not working with values a = %d, b = %d, expected = %d ", test.a, test.b, test.expected)
		}
	}

}
