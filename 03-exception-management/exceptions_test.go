package exceptions_test

import (
	"testing"

	"github.com/eljuanchosf/adventures-in-golang/03-exception-management"
)

func TestExceptions(t *testing.T) {
	_, err := exceptions.ThrowException()
	if err == nil {
		t.Error("ThrowException should have had returned and exception")
	}
}
