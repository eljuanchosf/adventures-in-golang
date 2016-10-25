package addition_test

import (
	"testing"

	"github.com/eljuanchosf/adventures-in-golang/01-testing-101"
)

func TestAddSomeNumbers(t *testing.T) {
	v := addition.AddSomeNumbers(1, 5)
	if v != 6 {
		t.Error("AddSomeNumbers doesn't work!")
	}
}
